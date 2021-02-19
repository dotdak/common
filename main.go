package main

import (
	"webserver-boilerplate/webserver"
)

func main() {
	s := webserver.NewWebserver()

	if err := s.StartServer(); err != nil {
		panic(err)
	}
}
