// package utils

// import (
// 	"flag"
// 	"tidy/multiple"
// )

// type Flags struct {
// 	DestDir     string
// 	InputFile   string
// 	RateLimit   int64
// 	URLs        []string
// }

// func ParseFlags() (*Flags, error) {
// 	flags := &Flags{}

// 	destDir := flag.String("P", ".", "Répertoire de destination pour le téléchargement")
// 	inputFile := flag.String("i", "", "Fichier texte contenant les liens à télécharger")
// 	rateLimitStr := flag.String("rate-limit", "0", "Limite de vitesse de téléchargement en KB/s (peut utiliser des unités k ou M)")

// 	flag.Parse()

// 	flags.DestDir = *destDir
// 	flags.InputFile = *inputFile

// 	rateLimit, err := multiple.ParseRateLimit(*rateLimitStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Convert int to int64
// 	flags.RateLimit = int64(rateLimit)

// 	if flag.NArg() > 0 {
// 		flags.URLs = append(flags.URLs, flag.Arg(0))
// 	}

// 	return flags, nil
// }

package utils

import (
	"flag"
	"tidy/multiple"
)

type Flags struct {
	DestDir      string
	InputFile    string
	RateLimit    int64
	URLs         []string
	Mirror       bool
	ConvertLinks bool
	Background bool
	RejectFiles  string
	ExcludeDirs  string
	OutputName string
}

func ParseFlags() (*Flags, error) {
	flags := &Flags{}

	// Standard flags
	destDir := flag.String("P", ".", "Répertoire de destination pour le téléchargement")
	inputFile := flag.String("i", "", "Fichier texte contenant les liens à télécharger")
	outputname := flag.String("O", "", "Fichier texte contenant les liens à télécharger")
	rateLimitStr := flag.String("rate-limit", "0", "Limite de vitesse de téléchargement en KB/s (peut utiliser des unités k ou M)")

	// Additional flags
	mirrorFlag := flag.Bool("mirror", false, "Mirror the website for offline use")
	backgroundFlag := flag.Bool("B", false, "Mirror the website for offline use")
	convertLinksFlag := flag.Bool("convert-links", false, "Convert links for offline viewing")
	rejectFlag := flag.String("R", "", "Comma-separated list of file extensions to reject (e.g., .jpg,.gif)")
	excludeFlag := flag.String("X", "", "Comma-separated list of directories to exclude (e.g., /assets,/css)")

	flag.Parse()

	// Set values in flags
	flags.DestDir = *destDir
	flags.OutputName = *outputname
	flags.Background = *backgroundFlag
	flags.InputFile = *inputFile
	flags.Mirror = *mirrorFlag
	flags.ConvertLinks = *convertLinksFlag
	flags.RejectFiles = *rejectFlag
	flags.ExcludeDirs = *excludeFlag

	// Parse rate limit
	rateLimit, err := multiple.ParseRateLimit(*rateLimitStr)
	if err != nil {
		return nil, err
	}
	flags.RateLimit = int64(rateLimit)

	// Capture remaining command line arguments as URLs
	if flag.NArg() > 0 {
		flags.URLs = flag.Args()
	}

	return flags, nil
}
