package commands

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/service/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
)

type Commands struct {
	CreateProduct         CreateProductCommandHandler
	DeleteProductByID     DeleteProductByIDCommandHandler
	DeactivateProductByID DeactivateProductByIDCommandHandler
	UpdateProductByID     UpdateProductByIDCommandHandler
}

func NewCommands(log logger.Logger, kafkaGateway *product.ProductKafkaGateway, v *validator.Validate, rs *grpc.ReaderService) *Commands {
	createHandler := NewCreateProductHandler(log, kafkaGateway, v)
	deleteHandler := NewDeleteProductByIDHandler(log, kafkaGateway, v, rs)
	deactivateHandler := NewDeactivateProductByIDHandler(log, kafkaGateway, v, rs)
	updateHandler := NewUpdateProductByIDHandler(log, kafkaGateway, v, rs)

	return &Commands{CreateProduct: createHandler, DeleteProductByID: deleteHandler, DeactivateProductByID: deactivateHandler, UpdateProductByID: updateHandler}
}
