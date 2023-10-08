package utils

import (
	"logSaver/pkg/model"
)

func SplitTime(logData model.LogData) (string, string, string) {
	createdTime := logData.GetCreated().UTC().Format("2006-01-02 15:04:05.000")
	updatedTime := logData.GetUpdated().UTC().Format("2006-01-02 15:04:05.000")
	timezone := logData.GetCreated().UTC().Format("-07:00")

	return createdTime, updatedTime, timezone
}
