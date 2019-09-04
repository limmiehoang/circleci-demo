package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	AppName     string       `toml:"app_name"`
	Environment string       `toml:"environment"`
	Server      serverConfig `toml:"server"`
}

type serverConfig struct {
	Listen    string   `toml:"listen"`
	WhiteList []string `toml:"whitelist"`
}

func New(path string) (*Config, error) {
	var config Config

	md, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s: %s(%v)", path, err, md)
	}
	return &config, nil
}
