package server

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/config"
	categoryRestGateway "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/category"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/grpc"
	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	gateway "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product"
	grpc_reader "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/prometheus"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/rest"
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
	m         *prometheus.Metrics
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, echo: echo.New(), v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.m = prometheus.NewMetrics(s.cfg)

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

	rc, err := grpc.NewReaderServiceClient(ctx, s.cfg)
	if err != nil {
		s.log.Errorf("Error in connecting grpc reader service PORT ", err)
		return err
	}
	rs := grpc_reader.NewReaderService(rc)

	kafkaGateway := gateway.NewProductKafkaGateway(s.cfg, kafkaProducer)

	restClient := rest.NewRestClient(s.log)
	categoryRestGateway := categoryRestGateway.NewCategoryRestGateway(s.cfg, s.log, restClient)

	commands := commands.NewCommands(s.log, kafkaGateway, categoryRestGateway, s.v, rs)

	productHandlers := routes.NewProductsHandlers(s.echo.Group(s.cfg.Http.ProductsPath), s.log, *commands, s.m)
	productHandlers.MapRoutes()

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.log.Infof("Valkyrie Product Core on PORT: %s", s.cfg.Http.Port)

	s.runMetrics(cancel)

	<-ctx.Done()

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("echo.Server.Shutdown", err)
	}

	return nil
}
