package config

import (
	"regexp"

	"github.com/caarlos0/env/v6"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Service  string `env:"DB_SERVICE"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Server   string `env:"DB_SERVER"`
	Port     string `env:"DB_PORT"`
	IsTest   bool   `env:"IS_TEST"`
}

func ReadConfig() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
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
