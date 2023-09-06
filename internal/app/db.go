package app

import (
	"database/sql"
	"fmt"
	_ "fmt"
)

var dbParams = map[string]string{
	"service":  "XEPDB1",
	"username": "hr",
	"server":   "DESKTOP-DOVQPAO",
	"port":     "1521",
	"password": "hr",
}

type db struct {
}

func NewDB() db {
	return db{}
}

func (db) SetupDB() (*sql.DB, error) {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"]
	db, err := sql.Open("oracle", connectionString)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
