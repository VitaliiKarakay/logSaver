package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godror/godror"
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
	MessageId    string
}

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log", logHandler)

	r.Run(":8080")
}
func logHandler(c *gin.Context) {
	var logData Log
	if err := c.BindJSON(&logData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := SetupDB()

	_, err := db.Exec("INSERT INTO log (user_id, phone, action_id, action_title, action_type, message, sender, status, language, full_response, created, updated, message_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
		logData.UserID, logData.Phone, logData.ActionID, logData.ActionTitle, logData.ActionType,
		logData.Message, logData.Sender, logData.Status, logData.Language, logData.FullResponse,
		logData.Created, logData.Updated, logData.MessageId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	defer db.Close()

	c.JSON(http.StatusOK, gin.H{"message": "log saved"})
}

func SetupDB() (*sql.DB, error) {
	dsn := "hr/hr@//DESKTOP-DOVQPAO:1521/XEPDB1"
	db, err := sql.Open(godror.DriverName, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
