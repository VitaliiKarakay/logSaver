package model

import "time"

type LogData interface {
	GetCreated() time.Time
	GetUpdated() time.Time
}
