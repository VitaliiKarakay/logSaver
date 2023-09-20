package postgresstore

import (
	"database/sql"

	_ "github.com/lib/pq"

	"logSaver/pkg/config"
)

type DB struct {
	DB *sql.DB

	LogRepository
}

func New(conf *config.Config) (*DB, error) {
	connectionString := "user=" + conf.Username + " password=" + conf.Password + " host=" + conf.Server + " port=" + conf.Port + " dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
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
