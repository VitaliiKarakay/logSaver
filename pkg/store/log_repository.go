package store

import (
	"database/sql"
	"fmt"

	"logSaver/pkg/model"
)

var logTableName = "Log"

type LogRepository struct {
	Oracle *sql.DB
}

func newLogRepository(db *sql.DB) LogRepository {
	return LogRepository{
		Oracle: db,
	}
}

func (lr *LogRepository) Insert(logData model.Log) error {
	statement, err := lr.Oracle.Prepare(`INSERT INTO ` + logTableName + ` (user_id, phone, action_id, action_title, action_type, 
                 message, sender, status, language, full_response, created, updated, message_id, STATUSDELIVE, COST)
							   VALUES (:UserID, :Phone, :ActionID, :ActionTitle, :ActionType,
							           :Message, :Sender, :Status, :Language, :FullResponse, TO_TIMESTAMP_TZ(:Created, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'),
							           TO_TIMESTAMP_TZ(:Updated, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'), :MessageID, :StatusDelive, :Cost)`)
	if err != nil {
		return err
	}

	defer func() {
		statementErr := statement.Close()
		if statementErr != nil {
			fmt.Println(statementErr)
		}
	}()
	createdTime := logData.Created.UTC().Format("2006-01-02 15:04:05.000")
	updatedTime := logData.Updated.UTC().Format("2006-01-02 15:04:05.000")
	timezone := logData.Created.UTC().Format("-07:00")

	_, err = statement.Exec(logData.UserID, logData.Phone, logData.ActionID, logData.ActionTitle, logData.ActionType,
		logData.Message, logData.Sender, logData.Status, logData.Language, logData.FullResponse,
		createdTime+" "+timezone, updatedTime+" "+timezone, logData.MessageID, logData.StatusDelive, logData.Cost)

	if err != nil {
		return err
	}

	return nil
}

func (lr *LogRepository) Get(log *model.Log) (model.Log, error) {
	existLogData := model.Log{}
	statement, err := lr.Oracle.Prepare(`SELECT USER_ID, PHONE, ACTION_ID, ACTION_TITLE, ACTION_TYPE,
       MESSAGE, SENDER, STATUS, LANGUAGE, FULL_RESPONSE, CREATED, UPDATED, MESSAGE_ID, STATUSDELIVE, COST FROM ` +
		logTableName + ` WHERE MESSAGE_ID = :MessageID AND PHONE = :Phone AND SENDER = :Sender`)
	if err != nil {
		return existLogData, err
	}

	defer func() {
		statementErr := statement.Close()
		if statementErr != nil {
			fmt.Println(statementErr)
		}
	}()

	result := statement.QueryRow(log.MessageID, log.Phone, log.Sender)
	err = result.Scan(
		&existLogData.UserID,
		&existLogData.Phone,
		&existLogData.ActionID,
		&existLogData.ActionTitle,
		&existLogData.ActionType,
		&existLogData.Message,
		&existLogData.Sender,
		&existLogData.Status,
		&existLogData.Language,
		&existLogData.FullResponse,
		&existLogData.Created,
		&existLogData.Updated,
		&existLogData.MessageID,
		&existLogData.StatusDelive,
		&existLogData.Cost,
	)
	if err != nil {
		return existLogData, err
	}

	return existLogData, nil
}

func (lr *LogRepository) Update(logData model.Log) error {
	statement, err := lr.Oracle.Prepare(`UPDATE ` + logTableName + ` SET user_id = :UserID, phone = :Phone,
               action_id = :ActionID, action_title = :ActionTitle, action_type = :ActionType, message = :Message,
               sender = :Sender, status = :Status, language = :Language, full_response = :FullResponse,
               created = TO_TIMESTAMP_TZ(:Created, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'), updated = TO_TIMESTAMP_TZ(:Updated, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'), message_id = :MessageID, STATUSDELIVE = :StatusDelive,
               COST = :Cost WHERE MESSAGE_ID = :MessageID AND PHONE = :Phone AND SENDER = :Sender`)
	if err != nil {
		return err
	}

	defer func() {
		statementErr := statement.Close()
		if statementErr != nil {
			fmt.Println(statementErr)
		}
	}()
	createdTime := logData.Created.UTC().Format("2006-01-02 15:04:05.000")
	updatedTime := logData.Updated.UTC().Format("2006-01-02 15:04:05.000")
	timezone := logData.Created.UTC().Format("-07:00")

	_, err = statement.Exec(logData.UserID, logData.Phone, logData.ActionID, logData.ActionTitle, logData.ActionType,
		logData.Message, logData.Sender, logData.Status, logData.Language, logData.FullResponse,
		createdTime+" "+timezone, updatedTime+" "+timezone, logData.MessageID, logData.StatusDelive, logData.Cost)
	if err != nil {
		return err
	}

	return nil
}
