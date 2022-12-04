package main

import (
	"flag"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/server"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/labstack/gommon/log"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewAppLogger(cfg.Logger)
	logger.InitLogger()

	s := server.NewServer(logger, cfg)
	log.Fatal(s.Run())
}
