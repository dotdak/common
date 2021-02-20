package main

import (
	"net/http"

	"webserver-boilerplate/webserver"

	"github.com/labstack/echo/v4"
)

func NewAPIs(s *webserver.WebServer) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.GET("/", apiHello(s))

	return e
}

func apiHello(s *webserver.WebServer) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	}
}
