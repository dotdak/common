package webserver

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Webserver struct {
	config *Config
	logger *log.Logger
	e      *echo.Echo
	secret []byte
}

func NewWebserver(options ...func(*Webserver)) *Webserver {

	args := LoadArgs()

	config, err := LoadConfigFromFile(args.ConfigFile)
	if err != nil {
		panic(err)
	}

	s := &Webserver{
		config: config,
		logger: log.New(log.Writer(), "", log.LstdFlags),
	}

	for _, option := range options {
		option(s)
	}

	loadApi(s)

	s.e.HideBanner = !config.Web.Debug

	return s
}

func (s *Webserver) StartServer() error {
	return s.e.Start(s.config.Web.Addr)
}
