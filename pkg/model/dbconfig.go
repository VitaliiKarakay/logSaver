package model

import (
	"encoding/json"
	"os"
)

type Config struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Server   string `json:"server"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func ReadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
