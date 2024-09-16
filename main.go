package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tidy/download"
)

func main() {
	destDir := flag.String("P", ".", "Répertoire de destination pour le téléchargement")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [-P destination_dir] <URL>")
		os.Exit(1)
	}

	url := flag.Arg(0)

	downloader := download.NewDownloader(*destDir)

	err := downloader.Download(url)
	if err != nil {
		log.Fatalf("Erreur lors du téléchargement : %v", err)
	}

	fmt.Println("Téléchargement terminé avec succès.")
}
