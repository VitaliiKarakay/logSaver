package model

import "time"

type SMSLog struct {
	UserID       int `db:"USER_ID"`
	Phone        string
	ActionID     int    `db:"ACTION_ID"`
	ActionTitle  string `db:"ACTION_TITLE"`
	ActionType   string `db:"ACTION_TYPE"`
	Message      string
	Sender       string
	Status       string
	Language     string
	FullResponse string `db:"FULL_RESPONSE"`
	Created      time.Time
	Updated      time.Time
	MessageID    string `db:"MESSAGE_ID"`
	StatusDelive int
	Cost         float32
}

func (l *SMSLog) UpdateExistLog(newLogData *SMSLog) {
	l.Status = newLogData.Status
	l.StatusDelive = newLogData.StatusDelive
	l.Cost = newLogData.Cost
	l.Updated = newLogData.Updated
}

func (l *SMSLog) GetCreated() time.Time {
	return l.Created
}

func (l *SMSLog) GetUpdated() time.Time {
	return l.Updated
}
