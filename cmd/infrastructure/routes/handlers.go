package routes

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands"
	commandModel "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/metrics"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/logger"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type productsHandlers struct {
	group   *echo.Group
	log     logger.Logger
	c       commands.Commands
	metrics *metrics.Metrics
}

func NewProductsHandlers(
	group *echo.Group,
	log logger.Logger,
	c commands.Commands,
	metrics *metrics.Metrics,
) *productsHandlers {
	return &productsHandlers{group: group, log: log, c: c, metrics: metrics}
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
		cmd := commandModel.NewProductCommandToCreate(*req)

		p, err := h.c.CreateProduct.Handle(ctx, *cmd)
		if err != nil {
			h.log.WarnMsg("Create Product", err)
			h.metrics.ErrorHttpRequests.Inc()
			return c.JSON(http.StatusBadRequest, err)
		}

		h.log.Info("The product has been created. ID: ", p.ProductID)
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusCreated, p)
	}
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
			h.metrics.ErrorHttpRequests.Inc()
			return c.JSON(http.StatusBadRequest, err)
		}

		h.log.Infof("The product with ID %s has been deleted", productID)
		h.metrics.SuccessHttpRequests.Inc()
		return c.NoContent(http.StatusOK)
	}
}

func (h *productsHandlers) DeactivateProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		productID, err := uuid.FromString(c.Param("id"))
		if err != nil {
			h.log.WarnMsg("uuid.FromString", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := h.c.DeactivateProductByID.Handle(ctx, productID); err != nil {
			h.log.WarnMsg("DeactivateProductByID", err)
			h.metrics.ErrorHttpRequests.Inc()
			return c.JSON(http.StatusBadRequest, err)
		}

		h.log.Infof("The product with ID %s has been deactivated", productID)
		h.metrics.SuccessHttpRequests.Inc()
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

		productID, err := uuid.FromString(c.Param("id"))
		if err != nil {
			h.log.WarnMsg("Update Product", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		req.ProductID = productID
		cmd := commandModel.NewProductCommandToUpdate(*req)
		p, err := h.c.UpdateProductByID.Handle(ctx, *cmd)
		if err != nil {
			h.log.WarnMsg("Update Product", err)
			h.metrics.ErrorHttpRequests.Inc()
			return c.JSON(http.StatusBadRequest, err)
		}

		h.log.Infof("The product with ID %s has been updated", req.ProductID)
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, p)
	}
}
