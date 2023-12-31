package http

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"logSaver/pkg/store/mongostore"
	_ "logSaver/pkg/store/oraclestore"
	_ "logSaver/pkg/store/postgresstore"
)

func Run(db *mongostore.DB) error {
	handler := LogHandler{DB: db}
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log", handler.CreateLog)
	r.PUT("/log", handler.UpdateLog)

	return r.Run(":8080")
}
