package webserver

import (
	"github.com/dotdak/common/logger"

	"github.com/go-logr/logr"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	*echo.Echo
	config *Config
	logger logr.Logger
	secret []byte
}

type option func(*WebServer)

func NewWebServer(options ...option) *WebServer {
	s := &WebServer{
		logger: logger.LOG(),
		config: DefaultConfig(),
	}

	AddOptions(s, options...)

	s.Echo = LoadAPI(s)

	return s
}

func AddOptions(s *WebServer, options ...option) {
	for _, opt := range options {
		opt(s)
	}
}

func (s *WebServer) StartServer() error {
	return s.Echo.Start(s.config.Web.Addr)
}
