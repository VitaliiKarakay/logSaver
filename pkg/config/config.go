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

func ReadPostgresConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg.PostgresConfig)
	if err != nil {
		return nil, err
	}

	err = cfg.PostgresConfig.Validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ReadOracleConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg.OracleConfig)
	if err != nil {
		return nil, err
	}

	err = cfg.OracleConfig.Validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ReadMongoConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg.MongoConfig)
	if err != nil {
		return nil, err
	}

	err = cfg.MongoConfig.Validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
