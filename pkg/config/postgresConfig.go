package config

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var PostgresDriver = "postgres"

type PostgresConfig struct {
	Username string `env:"DB_USERNAME_POSTGRES"`
	Password string `env:"DB_PASSWORD_POSTGRES"`
	Server   string `env:"DB_SERVER_POSTGRES"`
	Port     string `env:"DB_PORT_POSTGRES"`
	DBName   string `env:"DB_NAME_POSTGRES"`
	IsTest   bool   `env:"IS_TEST"`
}

func (c PostgresConfig) GetConnectionString(config PostgresConfig) string {
	return "user=" + config.Username + " password=" + config.Password + " host=" + config.Server + " port=" +
		config.Port + " dbname=" + config.DBName + " sslmode=disable"
}

func (c PostgresConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
		validation.Field(&c.Server, validation.Required),
		validation.Field(&c.Port, validation.Match(regexp.MustCompile(`^\d+$`)).Error("Port should be a number")),
	)
}
