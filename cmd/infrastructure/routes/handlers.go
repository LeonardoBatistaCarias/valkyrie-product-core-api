package routes

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	commandModels "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/models"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type productsHandlers struct {
	group *echo.Group
	c     commands.Commands
	v     *validator.Validate
}

func NewProductsHandlers(
	group *echo.Group,
	c commands.Commands,
	v *validator.Validate,
) *productsHandlers {
	return &productsHandlers{group: group, c: c, v: v}
}

func (h *productsHandlers) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		req := &models.CreateProductRequest{}
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		req.ProductID = uuid.NewV4()
		createCommand := buildProductCommand(*req)

		p, err := h.c.CreateProduct.Handle(ctx, *createCommand)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, p)
	}
}

func buildProductCommand(req models.CreateProductRequest) *commandModels.ProductCommand {
	return commandModels.NewProductCommand(
		req.ProductID,
		req.Name,
		req.Description,
		req.Brand,
		req.Price,
		req.Quantity,
		req.CategoryID,
		nil,
		req.Active,
	)
}
