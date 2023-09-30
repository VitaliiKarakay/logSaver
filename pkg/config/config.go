package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	PostgresConfig
	OracleConfig
	MongoConfig
}

func ReadConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg.PostgresConfig)
	if err != nil {
		return nil, err
	}

	err = env.Parse(&cfg.OracleConfig)
	if err != nil {
		return nil, err
	}

	err = env.Parse(&cfg.MongoConfig)
	if err != nil {
		return nil, err
	}

	err = cfg.PostgresConfig.Validate()
	err = cfg.OracleConfig.Validate()
	err = cfg.MongoConfig.Validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
