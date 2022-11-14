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
)
