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
