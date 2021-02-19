package webserver

import "flag"

type Args struct {
	ConfigFile string
	Debug      bool
	Verbose    bool
}

func LoadArgs() *Args {
	configFile := flag.String("config", "config.conf", "Path to config file")
	debug := flag.Bool("debug", false, "Debug mode")
	verbose := flag.Bool("verb", false, "Verbose mode")
	flag.Parse()
	args := Args{
		ConfigFile: *configFile,
		Debug:      *debug,
		Verbose:    *verbose,
	}

	return &args
}
