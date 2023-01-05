package commands

import (
	"context"
	commandModel "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
	"time"
)

type UpdateProductByIDCommandHandler interface {
	Handle(ctx context.Context, cmd commandModel.ProductCommand) (*model.UpdateProductResponse, error)
}

type updateProductByIDHandler struct {
	log     logger.Logger
	gateway product.ProductGateway
	v       *validator.Validate
	rs      *grpc.ReaderService
}

func NewUpdateProductByIDHandler(log logger.Logger, productGateway product.ProductGateway, v *validator.Validate, rs *grpc.ReaderService) *updateProductByIDHandler {
	return &updateProductByIDHandler{log: log, gateway: productGateway, v: v, rs: rs}
}

func (c *updateProductByIDHandler) Handle(ctx context.Context, cmd commandModel.ProductCommand) (*model.UpdateProductResponse, error) {
	_, err := c.rs.GetProductByID(cmd.ProductID)
	if err != nil {
		c.log.Errorf("A product with ID: %s doesn't exist.", cmd.ProductID)
		return nil, err
	}

	p := product.NewProduct(cmd.ProductID, cmd.Name, cmd.Description, product.Brand(cmd.Brand), cmd.Price, cmd.Quantity, cmd.CategoryID, cmd.Active, time.Now())

	if err := c.v.StructCtx(ctx, p); err != nil {
		c.log.WarnMsg("validate", err)
		return nil, err
	}

	if err := c.gateway.UpdateProductByID(ctx, *p); err != nil {
		c.log.Errorf("Error in generating a novelty in UpdateProduct topic", err)
		return nil, err
	}

	return model.NewUpdateProductResponse(p.ProductID), nil
}
