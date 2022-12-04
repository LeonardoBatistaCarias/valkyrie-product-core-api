package constants

import "time"

const (
	WRITER_READ_TIMEOUT  = 10 * time.Second
	WRITER_WRITE_TIMEOUT = 10 * time.Second
	WRITER_REQUIRED_ACKS = -1
	WRITER_MAX_ATTEMPTS  = 3
)
