package model

import "time"

type Log struct {
	UserID       int
	Phone        string
	ActionID     int
	ActionTitle  string
	ActionType   string
	Message      string
	Sender       string
	Status       string
	Language     string
	FullResponse string
	Created      time.Time
	Updated      time.Time
	MessageID    string
	StatusDelive int
	Cost         float32
}
