package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	// Global config
	config *Config = DefaultConfig()
)

func C() *Config {
	return config
}

// load config from file
func LoadConfigFromFile(filepath string) error {
	_, err := toml.DecodeFile(filepath, config)
	if err != nil {
		return err
	}
	return nil
}

// load config from env
func LoadConfigFromEnv() error {
	// Complete the mapping of environment variables and Config objects.
	return env.Parse(config)
}
