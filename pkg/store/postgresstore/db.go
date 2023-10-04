package postgresstore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"logSaver/pkg/config"
)

type DB struct {
	DB *sql.DB

	LogRepository
}

func New(conf *config.Config) (*DB, error) {
	connectionString := conf.PostgresConfig.GetConnectionString()
	db, err := sql.Open(config.PostgresDriver, connectionString)
	if err != nil {
		return nil, err
	}
	conn := &DB{DB: db}
	if conf.PostgresConfig.IsTest {
		conn.setTableNames()
		conn.createTestTables()
	}
	conn.LogRepository = newLogRepository(db)

	return conn, nil
}

func (database *DB) CloseConnection() error {
	return database.DB.Close()
}

func (*DB) setTableNames() {
	SMSLogTableName = config.LogTest
}

func (database *DB) createTestTables() {
	query := `CREATE TABLE IF NOT EXISTS ` + SMSLogTableName + ` (
                                            id SERIAL PRIMARY KEY,
                                            user_id INTEGER,
                                            phone VARCHAR(20),
                                            action_id INTEGER,
                                            action_title VARCHAR(255),
                                            action_type VARCHAR(50),
                                            message VARCHAR(1000),
                                            sender VARCHAR(100),
                                            status VARCHAR(50),
                                            language VARCHAR(10),
                                            full_response VARCHAR(1000),
                                            created TIMESTAMP,
                                            updated TIMESTAMP,
                                            message_id VARCHAR(100),
                                            statusDelive INTEGER,
                                            cost NUMERIC(10,4),
                                            CONSTRAINT unique_message_phone_sender_3 UNIQUE (message_id, phone, sender)
);`
	_, err := database.DB.Exec(query)
	if err != nil {
		fmt.Println("createTestTables", err)
	}
}
