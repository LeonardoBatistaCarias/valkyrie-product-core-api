package config

import (
	"flag"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/rest"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/constants"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"os"

	kafkaConfig "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Product core API microservice config path")
}

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	Logger      *logger.Config      `mapstructure:"logger"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
	Http        Http                `mapstructure:"http"`
	Grpc        Grpc                `mapstructure:"grpc"`
	Rest        *rest.Config        `mapstructure:"rest"`
	Kafka       *kafkaConfig.Config `mapstructure:"kafka"`
	Prometheus  Prometheus          `mapstructure:"prometheus"`
}

type Http struct {
	Port                string `mapstructure:"port"`
	Development         bool   `mapstructure:"development"`
	BasePath            string `mapstructure:"basePath"`
	ProductsPath        string `mapstructure:"productsPath"`
	DebugHeaders        bool   `mapstructure:"debugHeaders"`
	HttpClientDebug     bool   `mapstructure:"httpClientDebug"`
	DebugErrorsResponse bool   `mapstructure:"debugErrorsResponse"`
}

type Grpc struct {
	ReaderServicePort string `mapstructure:"readerServicePort"`
}

type KafkaTopics struct {
	ProductCreate     kafkaConfig.TopicConfig `mapstructure:"productCreate"`
	ProductDelete     kafkaConfig.TopicConfig `mapstructure:"productDelete"`
	ProductUpdate     kafkaConfig.TopicConfig `mapstructure:"productUpdate"`
	ProductDeactivate kafkaConfig.TopicConfig `mapstructure:"productDeactivate"`
}

type Prometheus struct {
	ReadinessPath        string `mapstructure:"readinessPath"`
	LivenessPath         string `mapstructure:"livenessPath"`
	Port                 string `mapstructure:"port"`
	Pprof                string `mapstructure:"pprof"`
	PrometheusPath       string `mapstructure:"prometheusPath"`
	PrometheusPort       string `mapstructure:"prometheusPort"`
	CheckIntervalSeconds int    `mapstructure:"checkIntervalSeconds"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.CONFIG_PATH)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/%s", getwd, constants.BASE_CONFIG_PATH)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.DEFAULT_CONFIG_TYPE)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	httpPort := os.Getenv(constants.HTTP_PORT)
	if httpPort != "" {
		cfg.Http.Port = httpPort
	}

	kafkaBrokers := os.Getenv(constants.KAFKA_BROKERS)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	return cfg, nil
}
