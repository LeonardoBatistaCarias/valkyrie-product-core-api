package routes

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/create"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type productsHandlers struct {
	group           *echo.Group
	productCommands commands.ProductCommands
}

func NewProductsHandlers(
	group *echo.Group,
	productCommands commands.ProductCommands,
) *productsHandlers {
	return &productsHandlers{group: group, productCommands: productCommands}
}

func (h *productsHandlers) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		dto := &models.CreateProductDTO{}
		if err := c.Bind(dto); err != nil {
			return nil
		}

		createCommand := create.NewCreateProductCommand(dto.Name, dto.Description, dto.Brand, dto.Price, dto.Quantity, dto.CategoryID, nil, true)

		if err := h.productCommands.CreateProduct.Handle(ctx, *createCommand); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		return c.JSON(http.StatusCreated, nil)
	}
}
