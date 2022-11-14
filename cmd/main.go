package main

import (
	"flag"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/server"
	"github.com/labstack/gommon/log"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)
	log.Fatal(s.Run())
}
