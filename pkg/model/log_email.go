package model

import "time"

/*

userID: 2134496917,
email: 'test@test.com',
actionID: 324,
actionTitle: 'Good Action',
actionType: 'promoEmail'|'triggerEmail',
title: 'some title,
sendler: 'sendgrid',
status: 'send'|'delivered'|'open'|'click'|'unsubscribe',
fullResponce: '',
created: "2023-02-27T00:27:00.031Z",
updated: "2023-02-27T00:27:00.031Z",
uniqKey: '6774560000068401360004'

*/

type EmailLog struct {
	UserID       int
	Email        string
	ActionID     int
	ActionTitle  string
	ActionType   string
	Title        string
	Sendler      string
	Status       string
	FullResponse string
	Created      time.Time
	Updated      time.Time
	UniqueKey    string
}

func (l *EmailLog) UpdateExistLog(newLogData *EmailLog) {
	l.Status = newLogData.Status
}
