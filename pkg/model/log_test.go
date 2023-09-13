package model_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"logSaver/pkg/model"
)

func TestLog_UpdateExistLog(t *testing.T) {
	log := model.GetTestLog(t)
	newLog := model.Log{
		Status:       "Sent",
		StatusDelive: 2,
		Cost:         14.88,
		Updated:      time.Now(),
	}
	log.UpdateExistLog(&newLog)

	assert.Equal(t, log.Status, newLog.Status)
	assert.Equal(t, log.StatusDelive, newLog.StatusDelive)
	assert.Equal(t, log.Cost, newLog.Cost)
	assert.Equal(t, log.Updated, newLog.Updated)
}
