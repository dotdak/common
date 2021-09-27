package webserver

import (
	"github.com/dotdak/common/logger"

	"github.com/go-logr/logr"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	config *Config
	logger logr.Logger
	e      *echo.Echo
	secret []byte
}

type option func(*WebServer)

func NewWebServer(options ...option) *WebServer {
	s := &WebServer{
		logger: logger.LOG(),
		config: DefaultConfig(),
	}

	AddOptions(s, options...)

	s.e = LoadAPI(s)

	return s
}

func AddOptions(s *WebServer, options ...option) {
	for _, opt := range options {
		opt(s)
	}
}

func (s *WebServer) StartServer() error {
	return s.e.Start(s.config.Web.Addr)
}
