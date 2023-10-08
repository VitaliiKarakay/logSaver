package emaillog

import (
	"database/sql"
	"fmt"

	"logSaver/pkg/model"
	"logSaver/pkg/utils"
)

var TableName = "email_log"

type EmailRepository struct {
	DB *sql.DB
}

func (lr *EmailRepository) InsertEmailLog(logData model.EmailLog) error {
	query := `INSERT INTO ` + TableName + ` (user_id, 
												email, 
												action_id, 
												action_title,
												action_type,
												title, 
												sender, 
												status, 
												full_response, 
												created, 
												updated, 
												unique_key)
										VALUES (:UserID, 
												:Email, 
												:ActionID, 
												:ActionTitle, 
												:ActionType,
												:Title, 
												:Sender, 
												:Status, 
												:FullResponse,
												TO_TIMESTAMP_TZ(:Created, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'),
												TO_TIMESTAMP_TZ(:Updated, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'),
												:UniqueKey)`
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
	createdTime, updatedTime, timezone := utils.SplitTime(&logData)

	_, err = statement.Exec(
		logData.UserID,
		logData.Email,
		logData.ActionID,
		logData.ActionTitle,
		logData.ActionType,
		logData.Title,
		logData.Sender,
		logData.Status,
		logData.FullResponse,
		createdTime+" "+timezone,
		updatedTime+" "+timezone,
		logData.UniqueKey)

	if err != nil {
		return err
	}

	return nil
}

func (lr *EmailRepository) GetEmailLog(log *model.EmailLog) (model.EmailLog, error) {
	existLogData := model.EmailLog{}
	firstPart := "SELECT USER_ID, " +
		"EMAIL, " +
		"ACTION_ID, " +
		"ACTION_TITLE, " +
		"ACTION_TYPE," +
		"TITLE, " +
		"SENDER, " +
		"STATUS, " +
		"FULL_RESPONSE, " +
		"CREATED, " +
		"UPDATED, " +
		"UNIQUE_KEY FROM "
	lastPart := " WHERE " +
		"UNIQUE_KEY = :UniqueKey"
	resultQuery := firstPart + TableName + lastPart
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

	result := statement.QueryRow(log.UniqueKey)
	err = result.Scan(
		&existLogData.UserID,
		&existLogData.Email,
		&existLogData.ActionID,
		&existLogData.ActionTitle,
		&existLogData.ActionType,
		&existLogData.Title,
		&existLogData.Sender,
		&existLogData.Status,
		&existLogData.FullResponse,
		&existLogData.Created,
		&existLogData.Updated,
		&existLogData.UniqueKey,
	)
	if err != nil {
		return existLogData, err
	}

	return existLogData, nil
}

func (lr *EmailRepository) UpdateEmailLog(logData model.EmailLog) error {
	firstPart := "Update " + TableName
	lastPart := " SET user_id = :UserID, " +
		"email = :Email, " +
		"action_id = :ActionID, " +
		"action_title = :ActionTitle, " +
		"action_type = :ActionType, " +
		"title = :Title, " +
		"sender = :Sender, " +
		"status = :Status, " +
		"full_response = :FullResponse, " +
		"created = TO_TIMESTAMP_TZ(:Created, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'), " +
		"updated = TO_TIMESTAMP_TZ(:Updated, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'), " +
		"unique_key = :UniqueKey " +
		"WHERE unique_key = :UniqueKey"
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
	createdTime, updatedTime, timezone := utils.SplitTime(&logData)

	_, err = statement.Exec(
		logData.UserID,
		logData.Email,
		logData.ActionID,
		logData.ActionTitle,
		logData.ActionType,
		logData.Title,
		logData.Sender,
		logData.Status,
		logData.FullResponse,
		createdTime+" "+timezone,
		updatedTime+" "+timezone,
		logData.UniqueKey)
	if err != nil {
		return err
	}

	return nil
}
