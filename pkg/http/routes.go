package http

import (
	"github.com/gin-gonic/gin"

	"net/http"

	_ "logSaver/pkg/store/oraclestore"
	"logSaver/pkg/store/postgresstore"
	_ "logSaver/pkg/store/postgresstore"
)

func Run(db *postgresstore.DB) error {
	handler := LogHandler{DB: db}
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log/sms", handler.CreateSMSLog)
	r.PUT("/log/sms", handler.UpdateSMSLog)

	r.POST("/log/email", handler.CreateEmailLog)
	r.PUT("/log/email", handler.UpdateEmailLog)

	return r.Run(":8080")
}
