package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
	http2 "logSaver/pkg/http"
	"net/http"
)

func main() {
	handler := http2.LogHandler{}
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log", handler.HandleLog)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}
