package server

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/create"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/config"
	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	gateway "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"

	"github.com/pkg/errors"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	cfg       *config.Config
	echo      *echo.Echo
	kafkaConn *kafka.Conn
}

func NewServer(cfg *config.Config) *server {
	return &server{cfg: cfg, echo: echo.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	kafkaProducer := kafkaClient.NewProducer(s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close() // nolint: errcheck

	log.Printf("Starting Writer Kafka consumers")
	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close() // nolint: errcheck

	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}
	kafkaGateway := gateway.NewProductKafkaGateway(s.cfg, kafkaProducer)
	createProductHandler := create.NewCreateProductHandler(kafkaGateway)
	productCommands := commands.NewProductCommands(createProductHandler)
	productHandlers := routes.NewProductsHandlers(s.echo.Group(s.cfg.Http.ProductsPath), *productCommands)
	productHandlers.MapRoutes()

	go func() {
		if err := s.runHttpServer(); err != nil {
			log.Fatalf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	log.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)
	<-ctx.Done()

	return nil
}

const (
	maxHeaderBytes = 1 << 20
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

func (s *server) runHttpServer() error {
	s.mapRoutes()

	s.echo.Server.ReadTimeout = readTimeout
	s.echo.Server.WriteTimeout = writeTimeout
	s.echo.Server.MaxHeaderBytes = maxHeaderBytes

	return s.echo.Start(s.cfg.Http.Port)
}

func (s *server) mapRoutes() {
	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	s.echo.Use(middleware.BodyLimit(bodyLimit))
}
