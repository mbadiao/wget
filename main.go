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
	// Manually parse custom flags before flag.Parse()
	rejectFiles, excludeDirs := parseCustomFlags()

	// fmt.Println("rejecteddddddddddddddd", rejectFiles, excludeDirs)
	// Parse standard flags using utils.ParseFlags()
	flags, err := utils.ParseFlags()
	utils.CheckError(err, "Erreur lors du parsing des flags")

	// Add reject and exclude values from custom flags
	// flags.RejectFiles = append(flags.RejectFiles, rejectFiles...)
	// flags.ExcludeDirs = append(flags.ExcludeDirs, excludeDirs...)

	// Load URLs from the input file, if provided
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

	// Exit if no URLs are provided
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
		excluded = rejectFiles
		flags.ExcludeDirs = strings.Join(excluded, ",")
	}

	var namefile string
	if flags.OutputName != ""{
		namefile = flags.OutputName
	}else{
		namefile = ""
	}

	// Initialize the downloader with the rate limit and options
	downloader := download.NewDownloader(flags.DestDir, flags.RateLimit, flags.Mirror, flags.ConvertLinks, flags.RejectFiles, flags.ExcludeDirs, *flags)

	// Concurrent downloading with a wait group
	var wg sync.WaitGroup
	// fmt.Println("rejected", rejected, "excluded",excluded)
	for _, url := range flags.URLs {
		if !downloader.ShouldReject(url, rejected) || !downloader.ShouldExclude(url, excluded) && len(rejected) != 0 && len(excluded) != 0{
			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				// Download the file using the configured downloader
				err := downloader.Download(url, namefile)
				if err != nil {
					fmt.Printf("Error downloading %s: %v\n", url, err)
				}
			}(url)
		} else {
			fmt.Printf("url %s rejected \n", url)
		}
	}

	// Wait for all downloads to complete
	wg.Wait()
	// fmt.Println("All downloads completed.")
}
