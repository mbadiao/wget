package mirror

// import (
// 	"fmt"
// 	"strings"
// )

// // ExcludeExtensions vérifie si un fichier doit être exclu (par exemple les GIF)
// func ExcludeExtensions(url string, excluded []string) bool {
// 	for _, ext := range excluded {
// 		if strings.HasSuffix(url, ext) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // OptimizeDownload gère les exclusions et autres optimisations
// func OptimizeDownload(url, destDir string, excluded []string) error {
// 	if ExcludeExtensions(url, excluded) {
// 		fmt.Printf("Exclusion de %s\n", url)
// 		return nil
// 	}

// 	// Continuez avec le téléchargement si le fichier n'est pas exclu
// 	err := DownloadPage(url, destDir)
// 	if err != nil {
// 		return fmt.Errorf("erreur lors du téléchargement : %v", err)
// 	}

// 	return nil
// }
