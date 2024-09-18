package utils

import (
	"flag"
	"tidy/multiple"
)

type Flags struct {
	DestDir     string
	InputFile   string
	RateLimit   int64
	URLs        []string
}

func ParseFlags() (*Flags, error) {
	flags := &Flags{}

	destDir := flag.String("P", ".", "Répertoire de destination pour le téléchargement")
	inputFile := flag.String("i", "", "Fichier texte contenant les liens à télécharger")
	rateLimitStr := flag.String("rate-limit", "0", "Limite de vitesse de téléchargement en KB/s (peut utiliser des unités k ou M)")
	
	flag.Parse()

	flags.DestDir = *destDir
	flags.InputFile = *inputFile

	rateLimit, err := multiple.ParseRateLimit(*rateLimitStr)
	if err != nil {
		return nil, err
	}

	// Convert int to int64
	flags.RateLimit = int64(rateLimit)

	if flag.NArg() > 0 {
		flags.URLs = append(flags.URLs, flag.Arg(0))
	}

	return flags, nil
}
