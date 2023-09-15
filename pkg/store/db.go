package store

import (
	"database/sql"

	"logSaver/pkg/config"
)

type DB struct {
	Oracle *sql.DB

	LogRepository
}

func New(conf *config.Config) (*DB, error) {
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
