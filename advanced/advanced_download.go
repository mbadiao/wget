package advanced

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Download a file under a given name with the option -0
func DownloadUnderName(url, OutputName string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while fetching: %s", err)
	}

	defer response.Body.Close()

	file, err1 := os.Create(OutputName)

	if err1 != nil {
		return fmt.Errorf("error while creating the file: %s", err1)
	}

	_, err2 := io.Copy(file, response.Body)

	if err2 != nil {
		return fmt.Errorf("error while copying the file: %s", err2)
	}

	return nil
}

// Download a file in the background with the option -B
func DownloadinBackground() {

}
