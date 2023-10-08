package oraclestore

import (
	"database/sql"

	"logSaver/pkg/store/oraclestore/emaillog"
	"logSaver/pkg/store/oraclestore/smslog"
)

func newSmsLogRepository(db *sql.DB) smslog.SmsRepository {
	return smslog.SmsRepository{
		DB: db,
	}
}

func newEmailLogRepository(db *sql.DB) emaillog.EmailRepository {
	return emaillog.EmailRepository{
		DB: db,
	}
}
