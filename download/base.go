package download

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"tidy/multiple"
	"time"

	"golang.org/x/net/html"
)

type Downloader struct {
	DestDir      string
	RateLimit    int64
	Mirror       bool
	ConvertLinks bool
	RejectFiles  string
	ExcludeDirs  string
}

// type Options struct {
// 	Mirror       bool
// 	ConvertLinks bool
// 	RejectFiles  []string
// 	ExcludeDirs  []string
// }

// NewDownloader creates a new downloader with additional options

func NewDownloader(destDir string, rateLimit int64, mirror, convertLinks bool, RejectFiles string, ExcludeDirs string) *Downloader {
	return &Downloader{
		DestDir:      destDir,
		RateLimit:    rateLimit,
		Mirror:       mirror,
		ConvertLinks: convertLinks,
		RejectFiles:  RejectFiles,
		ExcludeDirs:  ExcludeDirs,
	}
}

func (d *Downloader) Download(url string) error {
	startTime := time.Now()

	if d.Mirror {
		return d.mirrorSite(url)
	}
	// Configuration du client HTTP : içi on configure l'en-tête User-agent du package HTTP
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	//configuration de l'En-tête User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	//envoi de la requête HTTP
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de la requête HTTP : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("statut HTTP non valide : %s", resp.Status)
	}

	fileName := filepath.Base(url)
	filePath := filepath.Join(d.DestDir, fileName)

	// Creation du fichier
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier : %v", err)
	}
	defer out.Close()

	progressBar := NewProgressBar(resp.ContentLength, fileName)

	var reader io.Reader = resp.Body
	if d.RateLimit > 0 {
		rateLimitReader := multiple.NewRateLimitedReader(resp.Body, int64(d.RateLimit*1024))
		reader = io.TeeReader(rateLimitReader, progressBar)
	} else {
		reader = io.TeeReader(resp.Body, progressBar)
	}

	_, err = io.Copy(out, reader)
	if err != nil {
		return fmt.Errorf("erreur lors de la copie du contenu : %v", err)
	}

	endTime := time.Now()

	fmt.Printf("\nDébut du téléchargement : %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Fin du téléchargement : %s\n", endTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Statut HTTP : %s\n", resp.Status)
	fmt.Printf("Taille du fichier : %d octets (%.2f MB)\n", resp.ContentLength, float64(resp.ContentLength)/1048576)
	fmt.Printf("Fichier sauvegardé : %s\n", filePath)

	return nil
}

// func (d *Downloader) DownloadMirror(urlStr string) error {
// 	if d.Options.Mirror {
// 		return d.mirrorSite(urlStr)
// 	}
// 	return fmt.Errorf("non-mirror download not implemented")
// }

func (d *Downloader) ShouldReject(path string, rejectedpath []string) bool {
	fmt.Println("fjdksqlm", d.RejectFiles)
	for _, ext := range rejectedpath {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

func (d *Downloader) ShouldExclude(path string, exclude []string) bool {
	for _, dir := range exclude {
		if strings.Contains(path, dir) {
			return true
		}
	}
	return false
}

func (d *Downloader) convertLinks(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			err := d.convertFileLinks(path, dir)
			if err != nil {
				return fmt.Errorf("failed to convert links in %s: %v", path, err)
			}
		}
		return nil
	})
}

func (d *Downloader) convertFileLinks(filePath, baseDir string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return err
	}

	var convert func(*html.Node)
	convert = func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attr string
			switch n.Data {
			case "a", "link":
				attr = "href"
			case "img", "script":
				attr = "src"
			}
			for i, a := range n.Attr {
				if a.Key == attr {
					relPath, err := filepath.Rel(filepath.Dir(filePath), filepath.Join(baseDir, a.Val))
					if err == nil {
						n.Attr[i].Val = relPath
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			convert(c)
		}
	}
	convert(doc)

	var buf strings.Builder
	err = html.Render(&buf, doc)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, []byte(buf.String()), 0644)
}

func (d *Downloader) downloadRobotsTxt(baseURL string, destDir string) error {
	robotsURL := baseURL + "/robots.txt"
	resp, err := http.Get(robotsURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		robotsPath := filepath.Join(destDir, "robots.txt")
		return os.WriteFile(robotsPath, content, 0644)
	}
	return nil
}

func resolveURL(base, ref string) string {
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	refURL, err := baseURL.Parse(ref)
	if err != nil {
		return ""
	}
	return refURL.String()
}

func (d *Downloader) mirrorSite(urlStr string) error {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	domain := baseURL.Hostname()
	siteDir := filepath.Join(d.DestDir, domain)
	err = os.MkdirAll(siteDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %v", siteDir, err)
	}

	var tempRejected []string
	if d.RejectFiles != "" {
		tempRejected = strings.Split(d.RejectFiles, ",")
	}

	var tempExcluded []string
	if d.ExcludeDirs != "" {
		tempExcluded = strings.Split(d.ExcludeDirs, ",")
	}

	err = d.downloadRobotsTxt(urlStr, siteDir)
	if err != nil {
		fmt.Printf("Error downloading robots.txt: %v\n", err)
	}

	err = os.MkdirAll(siteDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %v", siteDir, err)
	}

	visited := make(map[string]bool)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := d.downloadRecursive(urlStr, siteDir, visited, &wg, &mu, domain, tempRejected, tempExcluded)
		if err != nil {
			fmt.Printf("Error downloading %s: %v\n", urlStr, err)
		}
	}()

	wg.Wait()

	if d.ConvertLinks {
		err = d.convertLinks(siteDir)
		if err != nil {
			return fmt.Errorf("failed to convert links: %v", err)
		}
	}

	return nil
}

func (d *Downloader) downloadRecursive(urlStr, destDir string, visited map[string]bool, wg *sync.WaitGroup, mu *sync.Mutex, baseDomain string, rejected, excluded []string) error {
	mu.Lock()
	if visited[urlStr] {
		mu.Unlock()
		return nil
	}
	visited[urlStr] = true
	mu.Unlock()

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	if parsedURL.Hostname() != baseDomain {
		fmt.Printf("Skipping external URL: %s\n", urlStr)
		return nil
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %v", urlStr, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code %d for %s", resp.StatusCode, urlStr)
	}

	localPath, err := d.urlToFilePath(urlStr, destDir)
	if err != nil {
		return fmt.Errorf("failed to convert URL to file path: %v", err)
	}

	if d.ShouldReject(localPath, rejected) || d.ShouldExclude(localPath, excluded) {
		return nil
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read content from %s: %v", urlStr, err)
	}

	isDir := strings.HasSuffix(parsedURL.Path, "/") || strings.HasSuffix(parsedURL.Path, "sites")
	// fmt.Println("localpath", localPath, isDir)

	// isDir := d.CheckIfDirectory(parsedURL.Path, visited)

	fmt.Println("isDir", isDir, parsedURL.Path)

	if isDir {
		localPath = filepath.Join(localPath, "index.html")
	}

	dir := filepath.Dir(localPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory for %s: %v", localPath, err)
	}

	err = os.WriteFile(localPath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", localPath, err)
	}

	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "text/html") {
		links, err := d.extractLinks(bytes.NewReader(content), urlStr)
		if err != nil {
			return fmt.Errorf("failed to extract links from %s: %v", urlStr, err)
		}

		for _, link := range links {
			wg.Add(1)
			go func(link string) {
				defer wg.Done()
				err := d.downloadRecursive(link, destDir, visited, wg, mu, baseDomain, rejected, excluded)
				if err != nil {
					fmt.Printf("Error downloading %s: %v\n", link, err)
				}
			}(link)
		}
	}

	return nil
}

func (d *Downloader) CheckIfDirectory(path string, temp map[string]bool ) bool {
	response, err := http.Get(path)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	// Check if the Content-Type header indicates a directory listing
	contentType := response.Header.Get("Content-Type")
	if strings.Contains(contentType, "text/html") {
		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return false
		}

		// Check for common directory listing patterns
		bodyString := string(body)
		if strings.Contains(bodyString, "<a href=") &&
			(strings.Contains(bodyString, "Parent Directory") ||
				strings.Contains(bodyString, "Index of")) && !temp[path] {
			return true
		}
	}

	return false
}

func (d *Downloader) urlToFilePath(urlStr, baseDir string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// Get the path from the URL and ensure it has a file name
	path := u.Path
	// if path == "" || path == "/" {
	// 	path = "/index"
	// }

	// Ensure the path does not start with a leading slash
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	// Convert URL path to a valid file path
	fullPath := filepath.Join(baseDir, filepath.FromSlash(path))

	return fullPath, nil
}

func (d *Downloader) createDirForFile(filePath string) error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, os.ModePerm)
}

func (d *Downloader) extractLinks(body io.Reader, baseURL string) ([]string, error) {
	links := make([]string, 0)
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attr string
			switch n.Data {
			case "a", "link":
				attr = "href"
			case "img", "script":
				attr = "src"
			case "style":
				// Extract URLs from inline CSS
				if n.FirstChild != nil {
					cssLinks, _ := d.extractCSSLinks(strings.NewReader(n.FirstChild.Data), baseURL)
					links = append(links, cssLinks...)
				}
			}
			for _, a := range n.Attr {
				if a.Key == attr {
					link, err := url.Parse(a.Val)
					if err == nil {
						absURL := resolveURL(baseURL, link.String())
						links = append(links, absURL)
					}
					break
				}
				// Check for inline style attribute
				if a.Key == "style" {
					cssLinks, _ := d.extractCSSLinks(strings.NewReader(a.Val), baseURL)
					links = append(links, cssLinks...)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	extract(doc)

	return links, nil
}

func (d *Downloader) extractCSSLinks(css io.Reader, baseURL string) ([]string, error) {
	links := make([]string, 0)
	content, err := io.ReadAll(css)
	if err != nil {
		return nil, err
	}

	// Regex to match url() in CSS
	re := regexp.MustCompile(`url\(['"]?([^'"()]+)['"]?\)`)
	matches := re.FindAllSubmatch(content, -1)

	for _, match := range matches {
		if len(match) > 1 {
			link, err := url.Parse(string(match[1]))
			if err == nil {
				absURL := resolveURL(baseURL, link.String())
				links = append(links, absURL)
			}
		}
	}

	return links, nil
}
