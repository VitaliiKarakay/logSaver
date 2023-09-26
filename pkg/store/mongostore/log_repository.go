package mongostore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"logSaver/pkg/model"
)

var DBName = "mongodb"
var collectionName = "logs"

type LogRepository struct {
	DB *mongo.Client
}

func newLogRepository(db *mongo.Client) LogRepository {
	return LogRepository{
		DB: db,
	}
}

func (lr *LogRepository) Insert(logData model.Log) error {
	collection := lr.DB.Database(DBName).Collection(collectionName)

	_, err := collection.InsertOne(context.TODO(), logData)
	if err != nil {
		return err
	}

	return nil
}

func (lr *LogRepository) Get(log *model.Log) (model.Log, error) {
	collection := lr.DB.Database(DBName).Collection(collectionName)

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
	collection := lr.DB.Database(DBName).Collection(collectionName)

	filter := bson.M{
		"messageid": logData.MessageID,
		"phone":     logData.Phone,
		"sender":    logData.Sender,
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":       logData.UserID,
			"action_id":     logData.ActionID,
			"action_title":  logData.ActionTitle,
			"action_type":   logData.ActionType,
			"message":       logData.Message,
			"sender":        logData.Sender,
			"status":        logData.Status,
			"language":      logData.Language,
			"full_response": logData.FullResponse,
			"created":       logData.Created,
			"updated":       logData.Updated,
			"statusDelive":  logData.StatusDelive,
			"cost":          logData.Cost,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
