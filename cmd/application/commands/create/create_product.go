package create

import (
	"context"
	commandModels "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/models"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/models"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"log"
)

type CreateProductCommandHandler interface {
	Handle(ctx context.Context, cmd commandModels.ProductCommand) (*models.CreateProductResponse, error)
}

type createProductHandler struct {
	gateway product.ProductGateway
	v       *validator.Validate
}

func NewCreateProductHandler(productGateway product.ProductGateway) *createProductHandler {
	return &createProductHandler{gateway: productGateway, v: validator.New()}
}

func (c *createProductHandler) Handle(ctx context.Context, cmd commandModels.ProductCommand) (*models.CreateProductResponse, error) {
	p := product.NewProduct(cmd.ProductID, cmd.Name, cmd.Description, product.Brand(cmd.Brand), cmd.Price, cmd.Quantity, uuid.NewV4(), nil, cmd.Active)

	if err := c.v.StructCtx(ctx, p); err != nil {
		log.Printf("validate %v", err)
		return nil, nil
	}

	if err := c.gateway.Create(ctx, *p); err != nil {
		log.Printf("Error in generating a novelty in ProductCreate topic %v", err)
		return nil, err
	}

	return models.NewCreateProductResponse(p.ProductID), nil
}
