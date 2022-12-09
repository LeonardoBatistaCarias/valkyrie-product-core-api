package commands

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/service/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)

type DeleteProductCommandHandler interface {
	Handle(ctx context.Context, productID uuid.UUID) error
}

type deleteProductHandler struct {
	log     logger.Logger
	gateway product.ProductGateway
	v       *validator.Validate
	rs      grpc.ReaderService
}

func NewDeleteProductHandler(log logger.Logger, productGateway product.ProductGateway, v *validator.Validate, rs grpc.ReaderService) *deleteProductHandler {
	return &deleteProductHandler{log: log, gateway: productGateway, v: v, rs: rs}
}

func (c *deleteProductHandler) Handle(ctx context.Context, productID uuid.UUID) error {
	p, err := c.rs.GetProductByID(productID)
	if err != nil {
		c.log.Errorf("A product with ID: %s doesn't exist.", productID.String())
		return err
	}

	p.Deactivate()
	if err := c.gateway.DeleteProductByID(ctx, *p); err != nil {
		c.log.Errorf("Error in generating a novelty in ProductCreate topic", err)
		return err
	}

	return nil
}
