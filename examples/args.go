package main

import "flag"

type Args struct {
	ConfigFile string
	Debug      bool
	Verbose    bool
	Bolt       bool
}

func LoadArgs() *Args {
	configFile := flag.String("config", "config.conf", "Path to config file")
	bolt := flag.Bool("boltdb", true, "Use boltdb")
	debug := flag.Bool("debug", false, "Debug mode")
	verbose := flag.Bool("verb", false, "Verbose mode")
	flag.Parse()
	args := Args{
		ConfigFile: *configFile,
		Bolt:       *bolt,
		Debug:      *debug,
		Verbose:    *verbose,
	}

	return &args
}
