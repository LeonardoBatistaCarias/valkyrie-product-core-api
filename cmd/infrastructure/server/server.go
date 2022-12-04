package server

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/config"
	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	gateway "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/routes"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log       logger.Logger
	cfg       *config.Config
	echo      *echo.Echo
	kafkaConn *kafka.Conn
	v         *validator.Validate
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, echo: echo.New(), v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	kafkaProducer := kafkaClient.NewProducer(s.log, s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close()

	log.Printf("Starting Writer Kafka consumers")
	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close()

	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}
	kafkaGateway := gateway.NewProductKafkaGateway(s.cfg, kafkaProducer)
	commands := commands.NewCommands(s.log, kafkaGateway, s.v)

	productHandlers := routes.NewProductsHandlers(s.echo.Group(s.cfg.Http.ProductsPath), s.log, *commands)
	productHandlers.MapRoutes()

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.log.Infof("Valkyrie Product Core on PORT: %s", s.cfg.Http.Port)

	<-ctx.Done()

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("echo.Server.Shutdown", err)
	}

	return nil
}
