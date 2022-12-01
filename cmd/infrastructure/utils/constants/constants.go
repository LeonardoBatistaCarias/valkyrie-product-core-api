package constants

import "time"

const (
	// Config Constants
	CONFIG_PATH         = "CONFIG_PATH"
	DEFAULT_CONFIG_TYPE = "yaml"
	HTTP_PORT           = "HTTP_PORT"

	// Kafka Config Constants
	KAFKA_BROKERS        = "KAFKA_BROKERS"
	WRITER_READ_TIMEOUT  = 10 * time.Second
	WRITER_WRITE_TIMEOUT = 10 * time.Second
	WRITER_REQUIRED_ACKS = -1
	WRITER_MAX_ATTEMPTS  = 3

	// Base Config Path
	BASE_CONFIG_PATH = "cmd/infrastructure/config/config.yaml"

	// HTTP Server Constants
	MAX_HEADER_BYTES = 1 << 20
	BODY_LIMIT       = "2M"
	READ_TIMEOUT     = 15 * time.Second
	WRITE_TIMEOUT    = 15 * time.Second
	GZIP_LEVEL       = 5

	STACK_SIZE = 1 << 10 // 1 KB
)
