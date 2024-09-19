package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"tidy/download"
	"tidy/utils"
)

// parseCustomFlags handles --reject and --exclude flags
func parseCustomFlags() ([]string, []string) {
	rejectFiles := []string{}
	excludeDirs := []string{}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "--reject=") {
			rejectFiles = strings.Split(strings.TrimPrefix(arg, "--reject="), ",")
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			i--
		} else if strings.HasPrefix(arg, "--exclude=") {
			excludeDirs = strings.Split(strings.TrimPrefix(arg, "--exclude="), ",")
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			i--
		}
	}
	return rejectFiles, excludeDirs
}

func main() {
	rejectFiles, excludeDirs := parseCustomFlags()

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

	var rejected []string
	if flags.RejectFiles != "" && len(rejectFiles) == 0 {
		rejected = strings.Split(flags.RejectFiles, ",")
	} else if flags.RejectFiles == "" && len(rejectFiles) != 0 {
		rejected = rejectFiles
		flags.RejectFiles = strings.Join(rejected, ",")
	}

	var excluded []string
	if flags.ExcludeDirs != "" && len(excludeDirs) == 0 {
		excluded = strings.Split(flags.ExcludeDirs, ",")
	} else if flags.ExcludeDirs == "" && len(excludeDirs) != 0 {
		excluded = excludeDirs
		flags.ExcludeDirs = strings.Join(excluded, ",")
	}

	var namefile string
	if flags.OutputName != "" {
		namefile = flags.OutputName
	} else {
		namefile = ""
	}

	var rateLimit int64
	if flags.RateLimit != 0 {
		rateLimit = flags.RateLimit
	}

	downloader := download.NewDownloader(flags.DestDir, rateLimit, flags.Mirror, flags.ConvertLinks, flags.RejectFiles, flags.ExcludeDirs, *flags)

	var wg sync.WaitGroup
	for _, url := range flags.URLs {
		if !downloader.ShouldReject(url, rejected) || !downloader.ShouldExclude(url, excluded) && len(rejected) != 0 && len(excluded) != 0 {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				err := downloader.Download(url, namefile)
				if err != nil {
					fmt.Printf("Error downloading %s: %v\n", url, err)
				}
			}(url)
		} else {
			fmt.Printf("url %s rejected \n", url)
		}
	}

	wg.Wait()
}
