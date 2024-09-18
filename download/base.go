package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"tidy/multiple"
	"time"
)

type Downloader struct {
	DestDir   string
	RateLimit int
}

func NewDownloader(destDir string, rateLimit int) *Downloader {
	return &Downloader{DestDir: destDir, RateLimit: rateLimit}
}

func (d *Downloader) Download(url string) error {
	startTime := time.Now()

	// Configuration du client HTTP : içi on configure l'en-tête User-agent du package HTTP
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	//configuration de l'En-tête User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	//envoi de la requête HTTP
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de la requête HTTP : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("statut HTTP non valide : %s", resp.Status)
	}

	fileName := filepath.Base(url)
	filePath := filepath.Join(d.DestDir, fileName)

	// Creation du fichier
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier : %v", err)
	}
	defer out.Close()

	progressBar := NewProgressBar(resp.ContentLength, fileName)

	var reader io.Reader = resp.Body
	if d.RateLimit > 0 {
		rateLimitReader := multiple.NewRateLimitedReader(resp.Body, int64(d.RateLimit*1024))
		reader = io.TeeReader(rateLimitReader, progressBar)
	} else {
		reader = io.TeeReader(resp.Body, progressBar)
	}

	_, err = io.Copy(out, reader)
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
