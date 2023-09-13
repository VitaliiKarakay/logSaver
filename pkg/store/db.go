package store

import (
	"database/sql"
	"log"
	"logSaver/pkg/config"
)

type DB struct {
	Oracle *sql.DB

	LogRepository
}

func New(conf *config.Config) (*DB, error) {
	err := conf.Validate()
	if err != nil {
		log.Fatalf("Validation error: %v\n", err)
	}
	connectionString := "oracle://" + conf.Username + ":" + conf.Password + "@" +
		conf.Server + ":" + conf.Port + "/" + conf.Service
	db, err := sql.Open("oracle", connectionString)

	if err != nil {
		return nil, err
	}
	conn := &DB{Oracle: db}

	conn.LogRepository = newLogRepository(db)

	return conn, nil
}

func (database *DB) CloseConnection() error {
	return database.Oracle.Close()
}
