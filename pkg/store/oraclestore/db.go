package oraclestore

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"

	"logSaver/pkg/config"
	"logSaver/pkg/store/oraclestore/emaillog"
	"logSaver/pkg/store/oraclestore/smslog"
)

type DB struct {
	DB *sql.DB

	smslog.SmsRepository
	emaillog.EmailRepository
}

func New(conf *config.Config) (*DB, error) {
	connectionString := conf.OracleConfig.GetConnectionString()

	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		return nil, err
	}
	conn := &DB{DB: db}
	if conf.OracleConfig.IsTest {
		conn.setTableNames()
		conn.createTestTables()
	}
	conn.SmsRepository = newSmsLogRepository(db)
	conn.EmailRepository = newEmailLogRepository(db)

	return conn, nil
}

func (database *DB) CloseConnection() error {
	return database.DB.Close()
}

func (*DB) setTableNames() {
	SMSLogTableName = config.LogTest
}

func (database *DB) createTestTables() {
	query := `DECLARE
    table_exists NUMBER;
BEGIN
    SELECT COUNT(*) INTO table_exists FROM user_tables WHERE table_name = UPPER('` + SMSLogTableName + `');
    
    IF table_exists = 0 THEN
        EXECUTE IMMEDIATE '
        CREATE TABLE ` + SMSLogTableName + ` (
            id INTEGER GENERATED BY DEFAULT ON NULL AS IDENTITY, 
            user_id INTEGER, 
            phone VARCHAR2(20), 
            action_id INTEGER, 
            action_title VARCHAR2(255), 
            action_type VARCHAR2(50), 
            message VARCHAR2(1000), 
            sender VARCHAR2(100), 
            status VARCHAR2(50), 
            language VARCHAR2(10), 
            full_response VARCHAR2(1000), 
            created TIMESTAMP, 
            updated TIMESTAMP, 
            message_id VARCHAR2(100), 
            statusDelive INTEGER, 
            cost NUMBER(10,4), 
            CONSTRAINT logs_pk_3 PRIMARY KEY (id), 
            CONSTRAINT unique_message_phone_sender_3 UNIQUE (message_id, phone, sender)
        )';
    END IF;
END;`
	_, err := database.DB.Exec(query)
	if err != nil {
		fmt.Println("createTestTables", err)
	}
}
