package webserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func loadApi(s *Webserver) {
	e := echo.New()
	e.GET("/", apiHello(s))

	s.e = e
}

func apiHello(s *Webserver) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	}
}
