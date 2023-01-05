package commands

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)

type DeleteProductByIDCommandHandler interface {
	Handle(ctx context.Context, productID uuid.UUID) error
}

type deleteProductByIDHandler struct {
	log     logger.Logger
	gateway product.ProductGateway
	v       *validator.Validate
	rs      *grpc.ReaderService
}

func NewDeleteProductByIDHandler(log logger.Logger, productGateway product.ProductGateway, v *validator.Validate, rs *grpc.ReaderService) *deleteProductByIDHandler {
	return &deleteProductByIDHandler{log: log, gateway: productGateway, v: v, rs: rs}
}

func (c *deleteProductByIDHandler) Handle(ctx context.Context, productID uuid.UUID) error {
	p, err := c.rs.GetProductByID(productID)
	if err != nil {
		c.log.Errorf("A product with ID: %s doesn't exist.", productID.String())
		return err
	}

	p.Deactivate()
	if err := c.gateway.DeleteProductByID(ctx, productID); err != nil {
		c.log.Errorf("Error in generating a novelty in DeleteProduct topic", err)
		return err
	}

	return nil
}
