package commands

import (
	"context"
	commandModel "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/category"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/go-playground/validator"
	"time"
)

type CreateProductCommandHandler interface {
	Handle(ctx context.Context, cmd commandModel.ProductCommand) (*model.CreateProductResponse, error)
}

type createProductHandler struct {
	log logger.Logger
	pg  product.ProductGateway
	cg  category.CategoryGateway
	v   *validator.Validate
}

func NewCreateProductHandler(log logger.Logger, pg product.ProductGateway, cg category.CategoryGateway, v *validator.Validate) *createProductHandler {
	return &createProductHandler{log: log, pg: pg, cg: cg, v: v}
}

func (c *createProductHandler) Handle(ctx context.Context, cmd commandModel.ProductCommand) (*model.CreateProductResponse, error) {
	p := product.NewProduct(cmd.ProductID, cmd.Name, cmd.Description, product.Brand(cmd.Brand), cmd.Price, cmd.Quantity, cmd.CategoryID, cmd.Active, time.Now())

	if err := c.v.StructCtx(ctx, p); err != nil {
		c.log.WarnMsg("validate", err)
		return nil, err
	}

	if _, err := c.cg.FindCategoryByID(ctx, p.CategoryID.String()); err != nil {
		c.log.WarnMsg("FindCategoryByID", err)
		return nil, err
	}

	if err := c.pg.CreateProduct(ctx, *p); err != nil {
		c.log.Errorf("Error in generating a novelty in ProductCreate topic", err)
		return nil, err
	}

	return model.NewCreateProductResponse(p.ProductID), nil
}
