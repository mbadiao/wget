package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Downloader struct {
	DestDir string
}

func NewDownloader(destDir string) *Downloader {
	return &Downloader{DestDir: destDir}
}

func (d *Downloader) Download(url string) error {
	startTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erreur lors de la requête HTTP : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("statut HTTP non valide : %s", resp.Status)
	}

	fileName := filepath.Base(url)
	filePath := filepath.Join(d.DestDir, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier : %v", err)
	}
	defer out.Close()

	progressBar := NewProgressBar(resp.ContentLength)

	_, err = io.Copy(io.MultiWriter(out, progressBar), resp.Body)
	if err != nil {
		return fmt.Errorf("erreur lors de la copie du contenu : %v", err)
	}

	endTime := time.Now()

	fmt.Printf("\nDébut du téléchargement : %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Fin du téléchargement : %s\n", endTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Statut HTTP : %s\n", resp.Status)
	fmt.Printf("Taille du fichier : %d octets (%.2f MB)\n", resp.ContentLength, float64(resp.ContentLength)/1048576)
	fmt.Printf("Fichier sauvegardé : %s\n", filePath)

	return nil
}