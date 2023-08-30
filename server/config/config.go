package config

import (
	"github.com/caarlos0/env"
)

type ServerConfig struct {
	Environment string `env:"ENV,required"`
	Port        string `env:"PORT,required"`
}

func LoadEnvConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
