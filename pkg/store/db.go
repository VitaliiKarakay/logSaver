package store

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"logSaver/pkg/model"
)

var path string

func init() {
	path = "pkg/config/config.json"
}

type Db struct {
}

func NewDB() Db {
	return Db{}
}

func (Db) SetupDB() (*sql.DB, error) {
	config, err := model.ReadConfig(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	connectionString := "oracle://" + config.Username + ":" + config.Password + "@" + config.Server + ":" + config.Port + "/" + config.Service
	db, err := sql.Open("oracle", connectionString)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
