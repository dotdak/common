package main

import (
	"github.com/dotdak/common/logger"
	"github.com/dotdak/common/webserver"
)

func main() {

	// args := LoadArgs()

	s := webserver.NewWebServer(
		// webserver.SetConfig(args.ConfigFile),
		webserver.SetSecret(nil),
	)

	logger.DEBUG().Info("This is test server", "server", s)
	// if args.Bolt {
	// 	webserver.AddOptions(s, webserver.LoadBoltDB())
	// }

	if err := s.StartServer(); err != nil {
		panic(err)
	}
}
