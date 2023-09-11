package http

import (
	"github.com/gin-gonic/gin"
	"logSaver/pkg/store"
	"net/http"
)

func Run(db *store.DB) error {
	handler := LogHandler{DB: db}
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log", handler.HandleLog)

	if err := r.Run(":8080"); err != nil {
		return err
	}

	return nil
}
