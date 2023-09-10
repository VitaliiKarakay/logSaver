package store

import (
	"database/sql"
	"fmt"
	"logSaver/pkg/model"
)

var path = "pkg/config/config.json"

type DB struct {
}

func NewDB() DB {
	return DB{}
}

func (DB) SetupDB() (*sql.DB, error) {
	config, err := model.ReadConfig(path)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}
	connectionString := "oracle://" + config.Username + ":" + config.Password + "@" +
		config.Server + ":" + config.Port + "/" + config.Service
	db, err := sql.Open("oracle", connectionString)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	return db, nil
}
