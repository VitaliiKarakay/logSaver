package model

import "time"

type Log struct {
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

func (l *Log) UpdateExistLog(newLogData *Log) {
	l.Status = newLogData.Status
	l.StatusDelive = newLogData.StatusDelive
	l.Cost = newLogData.Cost
	l.Updated = newLogData.Updated
}
