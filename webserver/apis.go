package webserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadAPI(s *WebServer) *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", apiHealthCheck(s))
	e.GET("/health", apiHealthCheck(s))

	return e
}

func apiHealthCheck(s *WebServer) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Status OK")
	}
}
