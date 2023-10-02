package postgresstore

import (
	"database/sql"
	"fmt"

	"logSaver/pkg/model"
)

var logTableName = "Log"

type LogRepository struct {
	DB *sql.DB
}

func newLogRepository(db *sql.DB) LogRepository {
	return LogRepository{
		DB: db,
	}
}

func (lr *LogRepository) Insert(logData model.Log) error {
	query := `INSERT INTO ` + logTableName + ` (user_id, 
												phone, 
												action_id, 
												action_title, 
												action_type, 
												message, 
												sender, 
												status, 
												language, 
												full_response, 
												created, 
												updated, 
												message_id, 
												statusdelive, 
												cost)
           VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
	statement, err := lr.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func() {
		statementErr := statement.Close()
		if statementErr != nil {
			fmt.Println(statementErr)
		}
	}()
	createdTime, updatedTime, timezone := splitTime(&logData)

	_, err = statement.Exec(
		logData.UserID,
		logData.Phone,
		logData.ActionID,
		logData.ActionTitle,
		logData.ActionType,
		logData.Message,
		logData.Sender,
		logData.Status,
		logData.Language,
		logData.FullResponse,
		createdTime+" "+timezone,
		updatedTime+" "+timezone,
		logData.MessageID,
		logData.StatusDelive,
		logData.Cost)

	if err != nil {
		return err
	}

	return nil
}

func (lr *LogRepository) Get(log *model.Log) (model.Log, error) {
	existLogData := model.Log{}
	firstPart := "SELECT USER_ID, " +
		"PHONE, " +
		"ACTION_ID, " +
		"ACTION_TITLE, " +
		"ACTION_TYPE," +
		"MESSAGE, " +
		"SENDER, " +
		"STATUS, " +
		"LANGUAGE, " +
		"FULL_RESPONSE, " +
		"CREATED, " +
		"UPDATED, " +
		"MESSAGE_ID, " +
		"STATUSDELIVE, " +
		"COST" +
		" FROM " + logTableName
	lastPart := " WHERE MESSAGE_ID = $1 AND PHONE = $2 AND SENDER = $3"
	resultQuery := firstPart + lastPart
	statement, err := lr.DB.Prepare(resultQuery)
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
	firstPart := "UPDATE " + logTableName
	lastPart := " SET user_id = $1, " +
		"phone = $2, " +
		"action_id = $3, " +
		"action_title = $4," +
		" action_type = $5, " +
		"message = $6, " +
		"sender = $7, " +
		"status = $8," +
		" language = $9, " +
		"full_response = $10," +
		" created = $11," +
		" updated = $12," +
		" message_id = $13, " +
		"STATUSDELIVE = $14, " +
		"COST = $15 " +
		"WHERE MESSAGE_ID = $13 AND PHONE = $2 AND SENDER = $7"
	resultQuery := firstPart + lastPart
	statement, err := lr.DB.Prepare(resultQuery)
	if err != nil {
		return err
	}

	defer func() {
		statementErr := statement.Close()
		if statementErr != nil {
			fmt.Println(statementErr)
		}
	}()
	createdTime, updatedTime, timezone := splitTime(&logData)

	_, err = statement.Exec(
		logData.UserID,
		logData.Phone,
		logData.ActionID,
		logData.ActionTitle,
		logData.ActionType,
		logData.Message,
		logData.Sender,
		logData.Status,
		logData.Language,
		logData.FullResponse,
		createdTime+" "+timezone,
		updatedTime+" "+timezone,
		logData.MessageID,
		logData.StatusDelive,
		logData.Cost)
	if err != nil {
		return err
	}

	return nil
}

func splitTime(logData *model.Log) (string, string, string) {
	createdTime := logData.Created.UTC().Format("2006-01-02 15:04:05.000")
	updatedTime := logData.Updated.UTC().Format("2006-01-02 15:04:05.000")
	timezone := logData.Created.UTC().Format("-07:00")

	return createdTime, updatedTime, timezone
}
