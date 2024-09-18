package mirror

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ConvertLinks(destDir string) error {
	fmt.Println("Converting links to local resources...")
	// Traverse the directory and modify all HTML files
	err := filepath.Walk(destDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".html" {
			err := convertHTMLLinks(path)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to convert links: %v", err)
	}

	fmt.Println("Link conversion completed.")
	return nil
}

// Helper function to modify HTML links
func convertHTMLLinks(filePath string) error {
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", filePath, err)
	}

	// Update all href/src attributes to point to local files
	content := string(input)
	// Example (simple replacement, extend as necessary):
	content = strings.ReplaceAll(content, "http://originalsite.com/", "./")

	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %v", filePath, err)
	}

	return nil
}
