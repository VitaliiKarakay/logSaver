package app

import (
	"database/sql"
	"fmt"
	_ "fmt"
)

var path string

func init() {
	path = "config/config.json"
}

type Db struct {
}

func NewDB() Db {
	return Db{}
}

func (Db) SetupDB() (*sql.DB, error) {
	config, err := ReadConfig(path)
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
