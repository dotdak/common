package webserver

import (
	"log"
)

func LoadConfig(path string) func(s *WebServer) {
	return func(s *WebServer) {
		config, err := LoadConfigFromFile(path)
		if err != nil {
			panic(err)
		}

		s.config = config
	}
}

func SetSecret(secret []byte) func(s *WebServer) {
	return func(s *WebServer) {
		s.secret = secret
	}
}

func DefaultLogger() func(s *WebServer) {
	return func(s *WebServer) {
		s.logger = log.New(log.Writer(), "", log.LstdFlags)
	}
}
