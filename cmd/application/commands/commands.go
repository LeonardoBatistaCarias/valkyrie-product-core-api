package commands

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/service/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
)

type Commands struct {
	CreateProduct CreateProductCommandHandler
	DeleteProduct DeleteProductCommandHandler
}

func NewCommands(log logger.Logger, kafkaGateway *product.ProductKafkaGateway, v *validator.Validate, rs grpc.ReaderService) *Commands {
	createHandler := NewCreateProductHandler(log, kafkaGateway, v)
	deleteHandler := NewDeleteProductHandler(log, kafkaGateway, v, rs)

	return &Commands{CreateProduct: createHandler, DeleteProduct: deleteHandler}
}
