package model

import "time"

type EmailLog struct {
	UserID       int
	Email        string
	ActionID     int
	ActionTitle  string
	ActionType   string
	Title        string
	Sender       string
	Status       string
	FullResponse string
	Created      time.Time
	Updated      time.Time
	UniqueKey    string
}

func (l *EmailLog) UpdateExistLog(newLogData *EmailLog) {
	l.Status = newLogData.Status
}

func (l *EmailLog) GetCreated() time.Time {
	return l.Created
}

func (l *EmailLog) GetUpdated() time.Time {
	return l.Updated
}
