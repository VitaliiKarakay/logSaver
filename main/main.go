package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
	"logSaver/internal/handlers"
	"net/http"
)

func main() {
	handler := handlers.LogHandler{}
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log", handler.HandleLog)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
