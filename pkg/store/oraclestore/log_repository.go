package oraclestore

import (
	"database/sql"
	"fmt"

	"logSaver/pkg/model"
)

var SMSLogTableName = "Log"
var EmailLogTableName = "email_log"

type LogRepository struct {
	DB *sql.DB
}

func newLogRepository(db *sql.DB) LogRepository {
	return LogRepository{
		DB: db,
	}
}

func (lr *LogRepository) InsertSMSLog(logData model.SMSLog) error {
	query := `INSERT INTO ` + SMSLogTableName + ` (user_id, 
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
												STATUSDELIVE, 
												COST)
										VALUES (:UserID, 
												:Phone, 
												:ActionID, 
												:ActionTitle, 
												:ActionType,
												:Message, 
												:Sender, 
												:Status, 
												:Language, 
												:FullResponse,
												TO_TIMESTAMP_TZ(:Created, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'),
												TO_TIMESTAMP_TZ(:Updated, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM'),
												:MessageID, :StatusDelive, :Cost)`
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

func (lr *LogRepository) GetSMSLog(log *model.SMSLog) (model.SMSLog, error) {
	existLogData := model.SMSLog{}
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
		"COST FROM "
	lastPart := " WHERE " +
		"MESSAGE_ID = :MessageID AND " +
		"PHONE = :Phone AND " +
		"SENDER = :Sender"
	resultQuery := firstPart + SMSLogTableName + lastPart
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

func (lr *LogRepository) UpdateSMSLog(logData model.SMSLog) error {
	firstPart := "UpdateSMSLog " + SMSLogTableName
	lastPart := " SET user_id = :UserID, " +
		"phone = :Phone, " +
		"action_id = :ActionID, " +
		"action_title = :ActionTitle," +
		" action_type = :ActionType, " +
		"message = :Message, " +
		"sender = :Sender, " +
		"status = :Status," +
		" language = :Language, " +
		"full_response = :FullResponse," +
		" created = TO_TIMESTAMP_TZ(:Created, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM')," +
		" updated = TO_TIMESTAMP_TZ(:Updated, 'YYYY-MM-DD HH24:MI:SS.FF TZH:TZM')," +
		" message_id = :MessageID, " +
		"STATUSDELIVE = :StatusDelive, " +
		"COST = :Cost " +
		"WHERE MESSAGE_ID = :MessageID AND PHONE = :Phone AND SENDER = :Sender"
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

func (lr *LogRepository) InsertEmailLog(logData model.EmailLog) error {
	query := `INSERT INTO ` + EmailLogTableName + ` (user_id, 
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
	createdTime, updatedTime, timezone := splitTime(&logData)

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

func splitTime(logData model.LogData) (string, string, string) {
	createdTime := logData.GetCreated().UTC().Format("2006-01-02 15:04:05.000")
	updatedTime := logData.GetUpdated().UTC().Format("2006-01-02 15:04:05.000")
	timezone := logData.GetCreated().UTC().Format("-07:00")

	return createdTime, updatedTime, timezone
}
