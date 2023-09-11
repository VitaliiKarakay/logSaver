package main

import (
	_ "github.com/sijms/go-ora/v2"
	"log"
	"logSaver/pkg/config"
	http2 "logSaver/pkg/http"
	"logSaver/pkg/store"
)

/*
{
    phone: '380953071221',
    sendler: 'intistele',
    status: 'success',
    statusDelive: 2,
    cost: 3.8351,
    updated: "2023-02-27T06:27:05.097Z",
    messageID: '6774560000068401360004'
  }

*/

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
	}

	err = db.CloseConnection()
	if err != nil {
		log.Fatal(err)
	}
}
