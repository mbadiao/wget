package multiple

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type RateLimitedReader struct {
	reader      io.Reader
	rateLimit   int64 // bytes per second
	tokenBucket int64
	lastRefill  time.Time
}

func NewRateLimitedReader(reader io.Reader, rateLimit int64) *RateLimitedReader {
	return &RateLimitedReader{
		reader:      reader,
		rateLimit:   rateLimit,
		tokenBucket: rateLimit, // Start with a full bucket
		lastRefill:  time.Now(),
	}
}

func (r *RateLimitedReader) Read(p []byte) (int, error) {
	now := time.Now()
	timePassed := now.Sub(r.lastRefill)
	r.lastRefill = now

	// Refill the token bucket
	r.tokenBucket += int64(timePassed.Seconds() * float64(r.rateLimit))
	if r.tokenBucket > r.rateLimit {
		r.tokenBucket = r.rateLimit // Cap at max capacity
	}

	// If bucket is empty, wait
	if r.tokenBucket <= 0 {
		time.Sleep(time.Second) // Wait for at least one second
		return 0, nil
	}

	// Limit the read size to available tokens
	readSize := int64(len(p))
	if readSize > r.tokenBucket {
		readSize = r.tokenBucket
	}

	n, err := r.reader.Read(p[:readSize])
	r.tokenBucket -= int64(n)

	return n, err
}

// lorsqu'on a comme flag rate-limit on verifie si l'unité est en kilobytes ou megabytes
func ParseRateLimit(rateLimitStr string) (int64, error) {
	var multiplier int64 = 1
	rateLimitStr = strings.TrimSpace(rateLimitStr)

	if strings.HasSuffix(rateLimitStr, "k") || strings.HasSuffix(rateLimitStr, "K") {
		multiplier = 1024
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "k")
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "K")
	} else if strings.HasSuffix(rateLimitStr, "M") || strings.HasSuffix(rateLimitStr, "m") {
		multiplier = 1024 * 1024
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "M")
		rateLimitStr = strings.TrimSuffix(rateLimitStr, "m")
	}

	rateLimit, err := strconv.ParseInt(rateLimitStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid rate-limit value: %s", rateLimitStr)
	}

	return rateLimit * multiplier, nil
}
