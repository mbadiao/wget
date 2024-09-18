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

/* 
package download

import "fmt"

type DownloadError struct {
	Message string
	Cause   error
}

func (e *DownloadError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func NewDownloadError(message string, cause error) *DownloadError {
	return &DownloadError{
		Message: message,
		Cause:   cause,
	}
}

*/