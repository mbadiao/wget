package utils

import (
	"fmt"
	"log"
	"os"
)

func CheckError(err error, message string) {
	if err != nil {
		log.Fatalf("%s : %v", message, err)
	}
}

func PrintUsageAndExit() {
	fmt.Println("Usage: go run main.go [-P destination_dir] [-i input_file] [--rate-limit rate] <URL>")
	os.Exit(1)
}