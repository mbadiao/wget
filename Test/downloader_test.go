package Test

// func TestDownload(t *testing.T) {
// 	// Create a test server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path != "/testfile.txt" {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("This is a test file content"))
// 	}))
// 	defer server.Close()

// 	// Create a temporary directory for the test
// 	tempDir, err := ioutil.TempDir("", "download_test")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp dir: %v", err)
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Create a downloader
// 	downloader := download.NewDownloader(tempDir, 0)

// 	// Test the Download function
// 	err = downloader.Download(server.URL + "/testfile.txt")
// 	if err != nil {
// 		t.Fatalf("Download failed: %v", err)
// 	}

// 	// Check if the file was downloaded correctly
// 	downloadedFile := filepath.Join(tempDir, "testfile.txt")
// 	content, err := ioutil.ReadFile(downloadedFile)
// 	if err != nil {
// 		t.Fatalf("Failed to read downloaded file: %v", err)
// 	}

// 	expectedContent := "This is a test file content"
// 	if string(content) != expectedContent {
// 		t.Errorf("Downloaded content does not match. Got %s, want %s", string(content), expectedContent)
// 	}
// }

// func TestDownloadError(t *testing.T) {
// 	// Create a test server that always returns 404
// 	server := httptest.NewServer(http.NotFoundHandler())
// 	defer server.Close()

// 	// Create a temporary directory for the test
// 	tempDir, err := ioutil.TempDir("", "download_test")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp dir: %v", err)
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Create a downloader
// 	downloader := download.NewDownloader(tempDir, 0)

// 	// Test the Download function with a non-existent URL
// 	err = downloader.Download(server.URL + "/nonexistent.txt")
// 	if err == nil {
// 		t.Fatalf("Expected an error, but got nil")
// 	}
// }
