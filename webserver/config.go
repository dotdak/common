package webserver

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type (
	WebConfig struct {
		Debug bool
		Addr  string
		TLS   bool
	}
	Config struct {
		Web WebConfig
	}
)

func LoadConfigFromFile(path string) (*Config, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
