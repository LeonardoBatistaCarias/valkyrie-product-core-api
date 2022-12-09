package constants

import "time"

const (
	BACKOFF_LINEAR  = 100 * time.Millisecond
	BACKOFF_RETRIES = 3
)
