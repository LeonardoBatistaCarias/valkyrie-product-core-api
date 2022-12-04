package routes

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	commandModels "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/models"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/models"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type productsHandlers struct {
	group *echo.Group
	log   logger.Logger
	c     commands.Commands
}

func NewProductsHandlers(
	group *echo.Group,
	log logger.Logger,
	c commands.Commands,
) *productsHandlers {
	return &productsHandlers{group: group, log: log, c: c}
}

func (h *productsHandlers) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		req := &models.CreateProductRequest{}
		if err := c.Bind(req); err != nil {
			h.log.WarnMsg("Bind", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		req.ProductID = uuid.NewV4()
		createCommand := buildProductCommand(*req)

		p, err := h.c.CreateProduct.Handle(ctx, *createCommand)
		if err != nil {
			h.log.WarnMsg("Create Product", err)
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
