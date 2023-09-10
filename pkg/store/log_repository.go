package store

import (
	"database/sql"
	"fmt"
	"logSaver/pkg/model"
)

type LogRepository struct {
	Oracle *sql.DB
}

func newLogRepository(db *sql.DB) LogRepository {
	return LogRepository{
		Oracle: db,
	}
}

func (lr *LogRepository) Insert(logData model.Log) error {
	statement, err := lr.Oracle.Prepare(`INSERT INTO LOG (user_id, phone, action_id, action_title, action_type, 
                 message, sender, status, language, full_response, created, updated, message_id)
							   VALUES (:UserID, :Phone, :ActionID, :ActionTitle, :ActionType,
							           :Message, :Sender, :Status, :Language, :FullResponse, :Created,
							           :Updated, :MessageID)`)
	if err != nil {
		return err
	}

	defer func() {
		statementErr := statement.Close()
		if statementErr != nil {
			fmt.Println(statementErr)
		}
	}()

	_, err = statement.Exec(logData.UserID, logData.Phone, logData.ActionID, logData.ActionTitle, logData.ActionType,
		logData.Message, logData.Sender, logData.Status, logData.Language, logData.FullResponse,
		logData.Created, logData.Updated, logData.MessageID)
	if err != nil {
		return err
	}

	return nil
}
