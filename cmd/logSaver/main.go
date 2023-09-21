package main

import (
	"log"

	"logSaver/pkg/config"
	"logSaver/pkg/http"
	"logSaver/pkg/store/oraclestore"
	_ "logSaver/pkg/store/postgresstore"
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

	db, err := oraclestore.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = db.CloseConnection()
		if err != nil {
			log.Println(err)
		}
	}()

	err = http.Run(db)
	if err != nil {
		log.Println(err)
	}
}
