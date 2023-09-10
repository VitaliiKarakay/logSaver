package http

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"logSaver/pkg/store"
	"net/http"
	"time"
)

type Log struct {
	UserID       int
	Phone        string
	ActionID     int
	ActionTitle  string
	ActionType   string
	Message      string
	Sender       string
	Status       string
	Language     string
	FullResponse string
	Created      time.Time
	Updated      time.Time
	MessageID    string
}

type LogHandler struct {
}

func (lh *LogHandler) HandleLog(context *gin.Context) {
	var logData Log
	if err := context.BindJSON(&logData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	newDB := store.NewDB()
	connection, err := newDB.SetupDB()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})

		return
	}
	defer func(connection *sql.DB) {
		connectionErr := connection.Close()
		if connectionErr != nil {
			fmt.Println(connectionErr)
		}
	}(connection)

	statement, err := connection.Prepare(`INSERT INTO LOG (user_id, phone, action_id, action_title, action_type, 
                 message, sender, status, language, full_response, created, updated, message_id)
							   VALUES (:UserID, :Phone, :ActionID, :ActionTitle, :ActionType,
							           :Message, :Sender, :Status, :Language, :FullResponse, :Created,
							           :Updated, :MessageID)`)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})

		return
	}
	defer func(statement *sql.Stmt) {
		statementErr := statement.Close()
		if statementErr != nil {
			fmt.Println(statementErr)
		}
	}(statement)

	_, err = statement.Exec(logData.UserID, logData.Phone, logData.ActionID, logData.ActionTitle, logData.ActionType,
		logData.Message, logData.Sender, logData.Status, logData.Language, logData.FullResponse,
		logData.Created, logData.Updated, logData.MessageID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log saved"})
}
