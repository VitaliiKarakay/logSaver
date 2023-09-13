package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Service  string `env:"DB_SERVICE"`
	Username string `env:"DB_USERNAME"`
	Server   string `env:"DB_SERVER"`
	Port     string `env:"DB_PORT"`
	Password string `env:"DB_PASSWORD"`
}

func ReadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %v\n", err)
	}
	return &cfg, nil
}
