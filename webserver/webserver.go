package webserver

import (
	"log"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	config *Config
	logger *log.Logger
	e      *echo.Echo
	secret []byte
}

func NewWebServer(options ...func(*WebServer)) *WebServer {

	s := &WebServer{}

	for _, option := range options {
		option(s)
	}

	s.e = LoadAPI(s)

	return s
}

func (s *WebServer) StartServer() error {
	return s.e.Start(s.config.Web.Addr)
}
