package oraclestore

import (
	"database/sql"

	_ "github.com/sijms/go-ora/v2"

	"logSaver/pkg/config"
)

type DB struct {
	DB *sql.DB

	LogRepository
}

func New(conf *config.Config) (*DB, error) {
	connectionString := "oracle://" + conf.Username + ":" + conf.Password + "@" +
		conf.Server + ":" + conf.Port + "/" + conf.Service

	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		return nil, err
	}
	conn := &DB{DB: db}
	if conf.IsTest {
		conn.setTableNames()
		//conn.createTestTables()
	}
	conn.LogRepository = newLogRepository(db)

	return conn, nil
}

func (database *DB) CloseConnection() error {
	return database.DB.Close()
}

func (*DB) setTableNames() {
	logTableName = config.LogTest
}
