package config

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var oracleDriver = "oracle"

type OracleConfig struct {
	Service  string `env:"DB_SERVICE_ORACLE"`
	Username string `env:"DB_USERNAME_ORACLE"`
	Password string `env:"DB_PASSWORD_ORACLE"`
	Server   string `env:"DB_SERVER_ORACLE"`
	Port     string `env:"DB_PORT_ORACLE"`
	IsTest   bool   `env:"IS_TEST"`
}

func (c OracleConfig) GetConnectionString(config OracleConfig) string {
	return oracleDriver + "://" + config.Username + ":" + config.Password + "@" +
		config.Server + ":" + config.Port + "/" + config.Service
}

func (c OracleConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Service, validation.Required),
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
		validation.Field(&c.Server, validation.Required),
		validation.Field(&c.Port, validation.Match(regexp.MustCompile(`^\d+$`)).Error("Port should be a number")),
	)
}
