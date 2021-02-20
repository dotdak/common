package main

import (
	"webserver-boilerplate/webserver"
)

func main() {

	args := LoadArgs()

	s := webserver.NewWebServer()

	e := NewAPIs(s)
	s = webserver.NewWebServer(
		webserver.LoadConfig(args.ConfigFile),
		webserver.DefaultLogger(),
		webserver.LoadAPI(e),
	)

	if err := s.StartServer(); err != nil {
		panic(err)
	}
}
