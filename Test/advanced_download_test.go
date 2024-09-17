package Test

import (
	"testing"
	"tidy/advanced"
)

func TestDownloadWithGivenName(T *testing.T) {
	err := advanced.DownloadUnderName("http://httpforever.com/", "first")
	if err != nil {
		T.Errorf("File c'ant be downloaded %s", err)
	}
}
