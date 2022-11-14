package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *productsHandlers) MapRoutes() {
	h.group.POST("", h.CreateProduct())
	h.group.Any("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Pong")
	})
}
