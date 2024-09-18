package download

import (
	"fmt"
	"strings"
	"time"
)

type ProgressBar struct {
	Total     int64
	Current   int64
	StartTime time.Time
	LastTime  time.Time
	FileName  string
}

func NewProgressBar(total int64, fileName string) *ProgressBar {
	return &ProgressBar{
		Total:     total,
		StartTime: time.Now(),
		LastTime:  time.Now(),
		FileName:  fileName,
	}
}

func (pb *ProgressBar) Write(p []byte) (int, error) {
	n := len(p)
	pb.Current += int64(n)
	pb.update()
	return n, nil
}

// cette fonction met à jour la barre de progression en affichant dans ce format : nom du fichier, pourcentage, barre, taille totale, vitesse, temps restant
func (pb *ProgressBar) update() {
	now := time.Now()
	if now.Sub(pb.LastTime) < time.Second/10 {
		return
	}
	pb.LastTime = now

	percent := int(float64(pb.Current) / float64(pb.Total) * 100)
	if pb.Current >= pb.Total {
		percent = 100
	}

	elapsed := now.Sub(pb.StartTime).Seconds()

	speed := float64(pb.Current) / elapsed / 1048576

	remaining := time.Duration(float64(pb.Total-pb.Current) / speed * float64(time.Second))

	totalMB := float64(pb.Total) / 1048576

	barWidth := 100
	bar := strings.Repeat("=", percent) + ">"
	if percent == 100 {
		bar = strings.Repeat("=", barWidth)
	}

	if pb.Current < pb.Total {
		percent++
	}

	fmt.Printf("\r%s  %3d%%[%-100s] %6.2fM  %.2fMB/s  ds %ds",
		pb.FileName,
		percent,
		bar,
		totalMB,
		speed,
		remaining.Round(time.Second))

	if pb.Current >= pb.Total {
		fmt.Println()
	}
}
