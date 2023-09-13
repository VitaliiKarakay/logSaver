package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var path = "pkg/config/config.json"

type Config struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Server   string `json:"server"`
	Port     string `json:"port"`
	Password string `json:"password"`
	IsTest   bool   `json:"isTest"`
}

func ReadConfig() (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
