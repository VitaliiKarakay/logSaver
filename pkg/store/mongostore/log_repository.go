package mongostore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"logSaver/pkg/model"
)

var DBName = "mongodb"
var collectionName = "logs"

type LogRepository struct {
	DB       *mongo.Client
	Database *mongo.Database
}

func newLogRepository(db *mongo.Client) LogRepository {
	database := db.Database(DBName)

	return LogRepository{
		DB:       db,
		Database: database,
	}
}

func (lr *LogRepository) Insert(logData model.Log) error {
	collection := lr.Database.Collection(collectionName)

	filter := bson.M{
		"messageid": logData.MessageID,
		"phone":     logData.Phone,
		"sender":    logData.Sender,
	}

	update := bson.M{
		"$set": logData, // Все поля из logData будут обновлены
	}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (lr *LogRepository) Get(log *model.Log) (model.Log, error) {
	collection := lr.Database.Collection(collectionName)

	filter := bson.M{
		"messageid": log.MessageID,
		"phone":     log.Phone,
		"sender":    log.Sender,
	}

	var existLogData model.Log
	err := collection.FindOne(context.TODO(), filter).Decode(&existLogData)
	if err != nil {
		return model.Log{}, err
	}

	return existLogData, nil
}

func (lr *LogRepository) Update(logData model.Log) error {
	collection := lr.Database.Collection(collectionName)

	filter := bson.M{
		"messageid": logData.MessageID,
		"phone":     logData.Phone,
		"sender":    logData.Sender,
	}

	update := bson.M{
		"$set": bson.M{
			"userid":       logData.UserID,
			"actionid":     logData.ActionID,
			"actiontitle":  logData.ActionTitle,
			"actiontype":   logData.ActionType,
			"message":      logData.Message,
			"sender":       logData.Sender,
			"status":       logData.Status,
			"language":     logData.Language,
			"fullresponse": logData.FullResponse,
			"created":      logData.Created,
			"updated":      logData.Updated,
			"statusdelive": logData.StatusDelive,
			"cost":         logData.Cost,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
