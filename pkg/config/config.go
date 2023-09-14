package config

import (
	"log"
	"regexp"

	"github.com/caarlos0/env/v6"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

type Config struct {
	Service  string `env:"DB_SERVICE"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Server   string `env:"DB_SERVER"`
	Port     string `env:"DB_PORT"`
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

	err = cfg.Validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Service, validation.Required),
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
		validation.Field(&c.Server, validation.Required),
		validation.Field(&c.Port, validation.Match(regexp.MustCompile(`^\d+$`)).Error("Post should ne a number")),
	)
}
