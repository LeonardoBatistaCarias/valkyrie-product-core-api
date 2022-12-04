package create

import (
	"context"
	commandModels "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/models"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/models"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
	"time"
)

type CreateProductCommandHandler interface {
	Handle(ctx context.Context, cmd commandModels.ProductCommand) (*models.CreateProductResponse, error)
}

type createProductHandler struct {
	log     logger.Logger
	gateway product.ProductGateway
	v       *validator.Validate
}

func NewCreateProductHandler(log logger.Logger, productGateway product.ProductGateway, v *validator.Validate) *createProductHandler {
	return &createProductHandler{log: log, gateway: productGateway, v: v}
}

func (c *createProductHandler) Handle(ctx context.Context, cmd commandModels.ProductCommand) (*models.CreateProductResponse, error) {
	p := product.NewProduct(cmd.ProductID, cmd.Name, cmd.Description, product.Brand(cmd.Brand), cmd.Price, cmd.Quantity, cmd.CategoryID, cmd.Active, time.Now())

	if err := c.v.StructCtx(ctx, p); err != nil {
		c.log.WarnMsg("validate", err)
		return nil, err
	}

	if err := c.gateway.Create(ctx, *p); err != nil {
		c.log.Errorf("Error in generating a novelty in ProductCreate topic", err)
		return nil, err
	}

	return models.NewCreateProductResponse(p.ProductID), nil
}
