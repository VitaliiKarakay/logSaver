package main

import (
	"database/sql"
	"fmt"
	"logSaver/internal/app"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
)

/*Необходимо создать сервис который будет принимать и сохранять логи от сервисов которые отправляют SMS и Email сообщения.
Какие данные будут приходить:
{
userID: 2134496917,
phone: '380953071221',
actionID: 324,
actionTitle: 'Good Action',
actionType: 'promoSMS',
message: 'some message',
sender: 'intistele',
status: 'success',
language: 'en',
fullResponce: 'sms_id: 6774560000068401360004',
created: "2023-02-27T00:27:00.031Z",
updated: "2023-02-27T00:27:00.031Z",
messageID: '6774560000068401360004'
}
БД: любая на твое усмотрение
Web: Gin */

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

func main() {

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log", logHandler)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func logHandler(c *gin.Context) {
	var logData Log
	if err := c.BindJSON(&logData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newDB := app.NewDB()
	db, err := newDB.SetupDB()
	if err != nil {
		fmt.Println(err)
	}
	statement, err := db.Prepare(`INSERT INTO LOG (user_id, phone, action_id, action_title, action_type, message, sender, status, language, full_response, created, updated, message_id)
							   VALUES (:UserID, :Phone, :ActionID, :ActionTitle, :ActionType, :Message, :Sender, :Status, :Language, :FullResponse, :Created, :Updated, :MessageID)`)
	_, err = statement.Exec(logData.UserID, logData.Phone, logData.ActionID, logData.ActionTitle, logData.ActionType,
		logData.Message, logData.Sender, logData.Status, logData.Language, logData.FullResponse,
		logData.Created, logData.Updated, logData.MessageID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	c.JSON(http.StatusOK, gin.H{"message": "log saved"})
}
