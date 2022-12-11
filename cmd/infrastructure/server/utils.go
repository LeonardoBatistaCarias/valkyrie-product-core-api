package server

import (
	"context"
	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"net"
	"strconv"
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaCon")
	}

	s.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	log.Info("kafka connected to brokers: %+v", brokers)

	return nil
}

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		log.Warnf("kafkaConn.Controller", err)
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	log.Infof("kafka controller uri: %s", controllerURI)

	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		log.Warnf("initKafkaTopics.DialContext", err)
		return
	}
	defer conn.Close()

	log.Infof("established new kafka controller connection: %s", controllerURI)

	productCreateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.ProductCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.ProductCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.ProductCreate.ReplicationFactor,
	}

	if err := conn.CreateTopics(
		productCreateTopic,
	); err != nil {
		log.Warnf("kafkaConn.CreateTopics", err)
		return
	}

	log.Infof("kafka topics created or already exists: %+v", []kafka.TopicConfig{productCreateTopic})
}

func (s *server) runMetrics(cancel context.CancelFunc) {
	metricsServer := echo.New()
	metricsServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         constants.STACK_SIZE,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	go func() {
		metricsServer.GET(s.cfg.Prometheus.PrometheusPath, echo.WrapHandler(promhttp.Handler()))
		s.log.Infof("Metrics server is running on port: %s", s.cfg.Prometheus.PrometheusPort)
		if err := metricsServer.Start(s.cfg.Prometheus.PrometheusPort); err != nil {
			s.log.Errorf("metricsServer.Start: %v", err)
			cancel()
		}
	}()
}
