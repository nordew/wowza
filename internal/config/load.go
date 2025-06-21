package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

func Load() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}