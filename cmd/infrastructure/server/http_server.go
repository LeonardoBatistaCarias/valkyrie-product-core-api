package server

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/constants"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"strings"
)

func (s *server) runHttpServer() error {
	s.mapRoutes()

	s.echo.Server.ReadTimeout = constants.READ_TIMEOUT
	s.echo.Server.WriteTimeout = constants.WRITE_TIMEOUT
	s.echo.Server.MaxHeaderBytes = constants.MAX_HEADER_BYTES

	return s.echo.Start(s.cfg.Http.Port)
}

func (s *server) mapRoutes() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Valkyrie Product Core"
	docs.SwaggerInfo.Description = "Valkyrie Product Core."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         constants.STACK_SIZE,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: constants.GZIP_LEVEL,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	s.echo.Use(middleware.BodyLimit(constants.BODY_LIMIT))
}
