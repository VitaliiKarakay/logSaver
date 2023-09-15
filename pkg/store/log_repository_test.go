package store_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"

	"logSaver/pkg/model"
)

func (s *StoreSuite) TestLogInsert() {
	log := model.CreateTestLog(s.T())

	err := s.Store.LogRepository.Insert(log)
	s.NoError(err)

	receivedLog, err := s.Store.LogRepository.Get(&log)
	fmt.Println(log)
	fmt.Println(receivedLog)
	assert.Equal(s.T(), log.UserID, receivedLog.UserID)
	assert.Equal(s.T(), log.Phone, receivedLog.Phone)
	assert.Equal(s.T(), log.ActionID, receivedLog.ActionID)
	assert.Equal(s.T(), log.ActionTitle, receivedLog.ActionTitle)
	assert.Equal(s.T(), log.ActionType, receivedLog.ActionType)
	assert.Equal(s.T(), log.Message, receivedLog.Message)
	assert.Equal(s.T(), log.Sender, receivedLog.Sender)
	assert.Equal(s.T(), log.Status, receivedLog.Status)
	assert.Equal(s.T(), log.Language, receivedLog.Language)
	assert.Equal(s.T(), log.FullResponse, receivedLog.FullResponse)
	assert.Equal(s.T(), log.Created.Unix(), receivedLog.Created.Unix())
	assert.Equal(s.T(), log.Created.Unix(), receivedLog.Created.Unix())
	assert.Equal(s.T(), log.MessageID, receivedLog.MessageID)
	assert.Equal(s.T(), log.StatusDelive, receivedLog.StatusDelive)
	assert.Equal(s.T(), log.Cost, receivedLog.Cost)
}
