package constants

import "time"

const (
	MAX_HEADER_BYTES = 1 << 20
	BODY_LIMIT       = "2M"
	READ_TIMEOUT     = 15 * time.Second
	WRITE_TIMEOUT    = 15 * time.Second
	GZIP_LEVEL       = 5

	STACK_SIZE = 1 << 10 // 1 KB
)
