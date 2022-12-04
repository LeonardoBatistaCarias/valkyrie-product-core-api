package kafka

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/constants"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

func NewWriter(brokers []string, errLogger kafka.Logger) *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: constants.WRITER_REQUIRED_ACKS,
		MaxAttempts:  constants.WRITER_MAX_ATTEMPTS,
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  constants.WRITER_READ_TIMEOUT,
		WriteTimeout: constants.WRITER_WRITE_TIMEOUT,
		Async:        false,
	}
	return w
}
