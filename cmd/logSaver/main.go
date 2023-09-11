package main

import (
	_ "github.com/sijms/go-ora/v2"
	"log"
	"logSaver/pkg/config"
	http2 "logSaver/pkg/http"
	"logSaver/pkg/store"
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

	err = http2.Run(db)
	if err != nil {
		log.Fatal(err)
	} //конец

	err = db.CloseConnection()
	if err != nil {
		log.Fatal(err)
	}
}
