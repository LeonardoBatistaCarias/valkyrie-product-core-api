package commands

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/service/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)

type DeactivateProductByIDCommandHandler interface {
	Handle(ctx context.Context, productID uuid.UUID) error
}

type deactivateProductByIDHandler struct {
	log     logger.Logger
	gateway product.ProductGateway
	v       *validator.Validate
	rs      grpc.ReaderService
}

func NewDeactivateProductByIDHandler(log logger.Logger, productGateway product.ProductGateway, v *validator.Validate, rs grpc.ReaderService) *deactivateProductByIDHandler {
	return &deactivateProductByIDHandler{log: log, gateway: productGateway, v: v, rs: rs}
}

func (c *deactivateProductByIDHandler) Handle(ctx context.Context, productID uuid.UUID) error {
	if _, err := c.rs.GetProductByID(productID); err != nil {
		c.log.Errorf("A product with ID: %s doesn't exist.", productID.String())
		return err
	}

	if err := c.gateway.DeactivateProductByID(ctx, productID); err != nil {
		c.log.Errorf("Error in generating a novelty in DeactivateProduct topic", err)
		return err
	}

	return nil
}
