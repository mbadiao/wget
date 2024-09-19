package download

import (
	"fmt"
	"os"
	"strings"
	"tidy/utils"
	"time"

	"golang.org/x/term"
)

type ProgressBar struct {
	Total      int64
	Current    int64
	StartTime  time.Time
	LastTime   time.Time
	FileName   string
	IsComplete bool
	RateLimit  float64 // New field to store the rate limit in bytes per second
	Flags      *utils.Flags
}

func NewProgressBar(total int64, fileName string, rateLimit float64, flags utils.Flags) *ProgressBar {
	return &ProgressBar{
		Total:      total,
		StartTime:  time.Now(),
		LastTime:   time.Now(),
		FileName:   fileName,
		IsComplete: false,
		RateLimit:  rateLimit,
		Flags:      &flags,
	}
}

func (pb *ProgressBar) Write(p []byte) (int, error) {
	n := len(p)
	pb.Current += int64(n)
	pb.update(false)
	return n, nil
}

func (pb *ProgressBar) update(force bool) {
	now := time.Now()

	if !force && now.Sub(pb.LastTime) < time.Second/10 {
		return
	}
	pb.LastTime = now

	percent := int(float64(pb.Current) / float64(pb.Total) * 100)
	if pb.Current >= pb.Total {
		percent = 100
		pb.IsComplete = true
	}

	elapsed := now.Sub(pb.StartTime).Seconds()
	speed := float64(pb.Current) / elapsed

	// Ensure the displayed speed never exceeds the rate limit
	if speed > pb.RateLimit {
		speed = pb.RateLimit
	}

	remaining := int(float64(pb.Total-pb.Current) / speed)
	if remaining < 0 {
		remaining = 0
	}

	totalMB := float64(pb.Total) / 1048576

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Fallback width if terminal size can't be determined
	}

	availableWidth := width - len(pb.FileName) - 35
	if availableWidth < 10 {
		availableWidth = 10
	}

	completedWidth := int(float64(availableWidth) * float64(percent) / 100)
	bar := strings.Repeat("=", completedWidth)
	if completedWidth < availableWidth {
		bar += ">"
		bar += strings.Repeat(" ", availableWidth-completedWidth-1)
	}

	var speedStr string
	if speed < 1024 {
		speedStr = fmt.Sprintf("%.2fB/s", speed)
	} else if speed < 1048576 {
		speedStr = fmt.Sprintf("%.2fKB/s", speed/1024)
	} else {
		speedStr = fmt.Sprintf("%.2fMB/s", speed/1048576)
	}

	var remainingStr string
	if remaining > 3600 {
		remainingStr = fmt.Sprintf("%dh %dm %ds", remaining/3600, (remaining%3600)/60, remaining%60)
	} else if remaining > 60 {
		remainingStr = fmt.Sprintf("%dm %ds", remaining/60, remaining%60)
	} else {
		remainingStr = fmt.Sprintf("%ds", remaining)
	}

	if !pb.Flags.Background {
		fmt.Printf("\r%s %3d%%[%s] %6.2fM %s %s",
			pb.FileName,
			percent,
			bar,
			totalMB,
			speedStr,
			remainingStr)

		if pb.IsComplete {
			fmt.Printf("\r%s 100%%[%s] %6.2fM %s 0s\n", pb.FileName, strings.Repeat("=", availableWidth)+">", totalMB, speedStr)
		}
	}
}

func (pb *ProgressBar) Finish() {
	pb.Current = pb.Total
	pb.update(true)
}
