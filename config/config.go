package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"18080"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
