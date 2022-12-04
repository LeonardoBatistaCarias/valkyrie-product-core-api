package commands

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/create"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
)

type Commands struct {
	CreateProduct create.CreateProductCommandHandler
}

func NewCommands(log logger.Logger, kafkaGateway *product.ProductKafkaGateway, v *validator.Validate) *Commands {
	createProductHandler := create.NewCreateProductHandler(log, kafkaGateway, v)

	return &Commands{CreateProduct: createProductHandler}
}
