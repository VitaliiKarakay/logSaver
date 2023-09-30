package mongostore

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"logSaver/pkg/config"
)

var logTableName = "Log"

type DB struct {
	DB *mongo.Client

	LogRepository
}

func New(conf *config.Config) (*DB, error) {
	connectionString := conf.MongoConfig.GetConnectionString(conf.MongoConfig)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	conn := &DB{DB: client}
	conn.LogRepository = NewLogRepository(client)

	return conn, nil
}

func (database *DB) CloseConnection() error {
	return database.DB.Disconnect(context.TODO())
}

func (*DB) setTableNames() {
	logTableName = config.LogTest
}
