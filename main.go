package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"tidy/download"
	"tidy/utils"
)

func main() {
	flags, err := utils.ParseFlags()
	utils.CheckError(err, "Erreur lors du parsing des flags")

	if flags.InputFile != "" {
		file, err := os.Open(flags.InputFile)
		utils.CheckError(err, "Erreur lors de l'ouverture du fichier")
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			flags.URLs = append(flags.URLs, scanner.Text())
		}

		utils.CheckError(scanner.Err(), "Erreur lors de la lecture du fichier")
	}

	if len(flags.URLs) == 0 {
		utils.PrintUsageAndExit()
	}

	// Convert int64 RateLimit to int for NewDownloader
	downloader := download.NewDownloader(flags.DestDir, int(flags.RateLimit))

	var wg sync.WaitGroup
	for _, url := range flags.URLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			err := downloader.Download(url)
			if err != nil {
				fmt.Printf("Erreur lors du téléchargement de %s : %v\n", url, err)
			} else {
				fmt.Printf("Téléchargement de %s terminé avec succès.\n", url)
			}
		}(url)
	}

	wg.Wait()
}
