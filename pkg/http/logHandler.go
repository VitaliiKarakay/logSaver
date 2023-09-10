package http

import (
	"fmt"
	"logSaver/pkg/model"
	"logSaver/pkg/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
}

func (lh *LogHandler) HandleLog(context *gin.Context) {
	logData := model.Log{}
	if err := context.BindJSON(&logData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	newDB := store.DB{}
	connection, err := newDB.SetupDB()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})

		return
	}
	defer func() {
		connectionErr := connection.Close()
		if connectionErr != nil {
			fmt.Println(connectionErr)
		}
	}()

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
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "log saved"})
}
