package oraclestore

import (
	"database/sql"

	"logSaver/pkg/store/oraclestore/emaillog"
	"logSaver/pkg/store/oraclestore/smslog"
)

var SMSLogTableName = "Log"

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
