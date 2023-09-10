package main

import (
	"fmt"
	"log"
	"logSaver/pkg/config"
	http2 "logSaver/pkg/http"
	"logSaver/pkg/store"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := store.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	handler := http2.LogHandler{DB: db}
	r := gin.Default() //должна запускаться в функции в http
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Message": "Log Saver"})
	})
	r.POST("/log", handler.HandleLog)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)

		return
	} //конец

	err = db.CloseConnection()
	if err != nil {
		log.Fatal(err)
	}
}
