package webserver

import "github.com/go-logr/logr"

func DefaultConfig() *Config {
	return &Config{
		Web: WebConfig{
			Addr: ":8080",
		},
		BoltDB: BoltDB{
			Path: "default.db",
		},
	}
}

func SetConfig(path string) func(s *WebServer) {
	return func(s *WebServer) {
		config, err := LoadConfigFromFile(path)
		if err != nil {
			s.config = DefaultConfig()
		}

		s.config = config
	}
}

func SetSecret(secret []byte) func(s *WebServer) {
	return func(s *WebServer) {
		s.secret = secret
	}
}

func SetLogger(logger logr.Logger) func(s *WebServer) {
	return func(s *WebServer) {
		s.logger = logger
	}
}

func SetAddr(addr string) func(s *WebServer) {
	return func(s *WebServer) {
		s.config.Web.Addr = addr
	}
}
