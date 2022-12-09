package routes

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	commandModel "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/model"
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

		req := &model.CreateProductRequest{}
		if err := c.Bind(req); err != nil {
			h.log.WarnMsg("Bind", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		req.ProductID = uuid.NewV4()
		cmd := buildProductCommand(*req)

		p, err := h.c.CreateProduct.Handle(ctx, *cmd)
		if err != nil {
			h.log.WarnMsg("Create Product", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, p)
	}
}

func buildProductCommand(req model.CreateProductRequest) *commandModel.ProductCommand {
	return commandModel.NewProductCommand(
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

func (h *productsHandlers) DeleteProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		productID, err := uuid.FromString(c.Param("id"))
		if err != nil {
			h.log.WarnMsg("uuid.FromString", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := h.c.DeleteProductByID.Handle(ctx, productID); err != nil {
			h.log.WarnMsg("DeleteProductById", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		h.log.Infof("The product with ID %s has been deleted", productID)
		return c.NoContent(http.StatusOK)
	}
}

func (h *productsHandlers) UpdateProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		req := &model.UpdateProductRequest{}
		if err := c.Bind(req); err != nil {
			h.log.WarnMsg("Bind", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		req.ProductID = uuid.NewV4()
		cmd := buildUpdateProductCommand(*req)

		p, err := h.c.CreateProduct.Handle(ctx, *cmd)
		if err != nil {
			h.log.WarnMsg("Create Product", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		h.log.Infof("The product with ID %s has been updated", req.ProductID)
		return c.JSON(http.StatusOK, p)
	}
}

func buildUpdateProductCommand(req model.UpdateProductRequest) *commandModel.ProductCommand {
	return commandModel.NewProductCommand(
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
