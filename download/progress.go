package download

import (
	"fmt"
	"time"
)

type ProgressBar struct {
	Total     int64
	Current   int64
	StartTime time.Time
	LastTime  time.Time
}

func NewProgressBar(total int64) *ProgressBar {
	return &ProgressBar{
		Total:     total,
		StartTime: time.Now(),
		LastTime:  time.Now(),
	}
}

func (pb *ProgressBar) Write(p []byte) (int, error) {
	n := len(p)
	pb.Current += int64(n)
	pb.update()
	return n, nil
}

func (pb *ProgressBar) update() {
	now := time.Now()
	if now.Sub(pb.LastTime) < time.Second/10 {
		return
	}
	pb.LastTime = now

	percent := float64(pb.Current) / float64(pb.Total) * 100
	elapsed := now.Sub(pb.StartTime).Seconds()
	speed := float64(pb.Current) / elapsed
	remaining := time.Duration(float64(pb.Total-pb.Current) / speed * float64(time.Second))

	fmt.Printf("\r%.2f%% | %.2f MB/%.2f MB | %.2f MB/s | %s restant",
		percent,
		float64(pb.Current)/1048576,
		float64(pb.Total)/1048576,
		speed/1048576,
		remaining.Round(time.Second))

	if pb.Current == pb.Total {
		fmt.Println()
	}
}
