package config

import (
	"flag"
	"fmt"
	kafkaConfig "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/utils/constants"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Product core API microservice config path")
}

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
	Http        Http                `mapstructure:"http"`
	Kafka       *kafkaConfig.Config `mapstructure:"kafka"`
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

type KafkaTopics struct {
	ProductCreate kafkaConfig.TopicConfig `mapstructure:"productCreate"`
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
			x := fmt.Sprintf("%s/cmd/infrastructure/config/config.yaml", getwd)
			println(x)
			configPath = fmt.Sprintf("%s/cmd/infrastructure/config/config.yaml", getwd)
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
