package mongostore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"logSaver/pkg/model"
)

const DBName = "mongodb"
const collectionSMSLog = "sms_logs"
const collectionEmailLog = "email_logs"

type LogRepository struct {
	DB       *mongo.Client
	Database *mongo.Database
}

func NewLogRepository(db *mongo.Client) LogRepository {
	database := db.Database(DBName)

	return LogRepository{
		DB:       db,
		Database: database,
	}
}

func (lr *LogRepository) InsertSMSLog(logData model.SMSLog) error {
	collection := lr.Database.Collection(collectionSMSLog)

	filter := bson.M{
		"messageid": logData.MessageID,
		"phone":     logData.Phone,
		"sender":    logData.Sender,
	}

	update := bson.M{
		"$set": logData,
	}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (lr *LogRepository) GetSMSLog(log *model.SMSLog) (model.SMSLog, error) {
	collection := lr.Database.Collection(collectionSMSLog)

	filter := bson.M{
		"messageid": log.MessageID,
		"phone":     log.Phone,
		"sender":    log.Sender,
	}

	existLogData := model.SMSLog{}
	err := collection.FindOne(context.TODO(), filter).Decode(&existLogData)
	if err != nil {
		return model.SMSLog{}, err
	}

	return existLogData, nil
}

func (lr *LogRepository) UpdateSMSLog(logData model.SMSLog) error {
	collection := lr.Database.Collection(collectionSMSLog)

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

func (lr *LogRepository) InsertEmailLog(data *model.EmailLog) error {
	collection := lr.Database.Collection(collectionEmailLog)

	filter := bson.M{
		"uniquekey": data.UniqueKey,
	}

	update := bson.M{
		"$set": data,
	}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (lr *LogRepository) GetEmailLog(data *model.EmailLog) (model.EmailLog, error) {
	collection := lr.Database.Collection(collectionEmailLog)

	filter := bson.M{
		"uniquekey": data.UniqueKey,
	}

	existLogData := model.EmailLog{}
	err := collection.FindOne(context.TODO(), filter).Decode(&existLogData)
	if err != nil {
		return model.EmailLog{}, err
	}

	return existLogData, nil
}

func (lr *LogRepository) UpdateEmailLog(data model.EmailLog) error {
	collectionEmailLog := lr.Database.Collection(collectionEmailLog)

	filter := bson.M{
		"uniqueKey": data.UniqueKey,
	}

	update := bson.M{
		"$set": bson.M{
			"userid":       data.UserID,
			"email":        data.Email,
			"actionid":     data.ActionID,
			"actiontitle":  data.ActionTitle,
			"actiontype":   data.ActionType,
			"title":        data.Title,
			"sendler":      data.Sendler,
			"status":       data.Status,
			"fullresponse": data.FullResponse,
			"created":      data.Created,
			"updated":      data.Updated,
			"uniquekey":    data.UniqueKey,
		},
	}

	_, err := collectionEmailLog.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (lr *LogRepository) DeleteAllLogs() error { // перенести в store_suite
	collection := lr.DB.Database(DBName).Collection(collectionSMSLog)

	_, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	return nil
}
