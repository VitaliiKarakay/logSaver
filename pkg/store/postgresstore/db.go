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
	connectionString := conf.PostgresConfig.GetConnectionString(conf.PostgresConfig)
	db, err := sql.Open(config.PostgresDriver, connectionString)
	if err != nil {
		return nil, err
	}
	conn := &DB{DB: db}
	if conf.PostgresConfig.IsTest {
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
