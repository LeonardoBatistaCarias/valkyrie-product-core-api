package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *productsHandlers) MapRoutes() {
	h.group.POST("", h.CreateProduct())
	h.group.DELETE("/:id", h.DeleteProductByID())
	h.group.PUT("/:id", h.UpdateProductByID())
	h.group.Any("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Pong")
	})
}
