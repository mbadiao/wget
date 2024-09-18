package multiple

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type RateLimitedReader struct {
	reader    io.Reader
	rateLimit int64
	ticker    *time.Ticker
}

func NewRateLimitedReader(reader io.Reader, rateLimit int64) *RateLimitedReader {
	return &RateLimitedReader{
		reader:    reader,
		rateLimit: rateLimit,
		ticker:    time.NewTicker(time.Second),
	}
}

func (r *RateLimitedReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if n > 0 {
		<-r.ticker.C
	}
	return n, err
}

// lorsqu'on a comme flag rate-limit on verifie si l'unité est en kilobytes ou megabytes
func ParseRateLimit(rateLimitStr string) (int, error) {
	// Vérifier si l'unité est en kilobytes ou megabytes
	if strings.HasSuffix(rateLimitStr, "k") || strings.HasSuffix(rateLimitStr, "K") {
		// Convertir les kilobytes en int
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "k")
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "K")
		rateLimit, err := strconv.Atoi(rateLimitStr)
		if err != nil {
			return 0, fmt.Errorf("invalid rate-limit value: %s", rateLimitStr)
		}
		return rateLimit, nil
	} else if strings.HasSuffix(rateLimitStr, "M") || strings.HasSuffix(rateLimitStr, "m") {
		// Convertir les megabytes en kilobytes
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "M")
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "m")
		rateLimit, err := strconv.Atoi(rateLimitStr)
		if err != nil {
			return 0, fmt.Errorf("invalid rate-limit value: %s", rateLimitStr)
		}
		return rateLimit * 1024, nil // 1 MB = 1024 KB
	} else {
		// Si pas d'unité, supposer que c'est en kilobytes
		rateLimit, err := strconv.Atoi(rateLimitStr)
		if err != nil {
			return 0, fmt.Errorf("invalid rate-limit value: %s", rateLimitStr)
		}
		return rateLimit, nil
	}
}
