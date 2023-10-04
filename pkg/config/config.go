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

func (cfg *Config) ReadDatabaseConfigs() error {
	err := cfg.readPostgresConfig()
	if err != nil {
		return err
	}

	err = cfg.readOracleConfig()
	if err != nil {
		return err
	}

	err = cfg.readMongoConfig()
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) readPostgresConfig() error {
	err := env.Parse(&cfg.PostgresConfig)
	if err != nil {
		return err
	}

	err = cfg.PostgresConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) readOracleConfig() error {
	err := env.Parse(&cfg.OracleConfig)
	if err != nil {
		return err
	}

	err = cfg.OracleConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) readMongoConfig() error {
	err := env.Parse(&cfg.MongoConfig)
	if err != nil {
		return err
	}

	err = cfg.MongoConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}
