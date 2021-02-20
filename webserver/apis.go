package webserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoadAPI(s *WebServer) *echo.Echo {

	e := echo.New()
	e.HideBanner = true

	e.GET("/", apiHello(s))

	return e
}

func apiHello(s *WebServer) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	}
}
