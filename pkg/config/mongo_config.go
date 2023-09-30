package config

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var mongoDriver = "mongodb"

type MongoConfig struct {
	Username string `env:"DB_USERNAME_MONGO"`
	Password string `env:"DB_PASSWORD_MONGO"`
	Server   string `env:"DB_SERVER_MONGO"`
	Port     string `env:"DB_PORT_MONGO"`
}

func (c MongoConfig) GetConnectionString() string {
	return mongoDriver + "://" + c.Username + ":" + c.Password + "@" +
		c.Server + ":" + c.Port
}

func (c MongoConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
		validation.Field(&c.Server, validation.Required),
		validation.Field(&c.Port, validation.Match(regexp.MustCompile(`^\d+$`)).Error("Port should be a number")),
	)
}
