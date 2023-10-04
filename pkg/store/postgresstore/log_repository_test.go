package postgresstore_test

import (
	"log"

	"github.com/stretchr/testify/assert"

	"logSaver/pkg/model"
)

func (s *StoreSuite) TestLogInsert() {
	newLog := model.CreateTestLog(s.T())

	err := s.Store.LogRepository.InsertSMSLog(newLog)
	s.NoError(err)

	err = s.Store.LogRepository.InsertSMSLog(newLog)
	if err != nil {
		log.Print(err)
	}
	s.NotNil(err)
}

func (s *StoreSuite) TestLogGet() {
	newLog := model.CreateTestLog(s.T())

	err := s.Store.LogRepository.InsertSMSLog(newLog)
	s.NoError(err)

	receivedLog, err := s.Store.LogRepository.GetSMSLog(&newLog)
	assert.Equal(s.T(), newLog.UserID, receivedLog.UserID)
	assert.Equal(s.T(), newLog.Phone, receivedLog.Phone)
	assert.Equal(s.T(), newLog.ActionID, receivedLog.ActionID)
	assert.Equal(s.T(), newLog.ActionTitle, receivedLog.ActionTitle)
	assert.Equal(s.T(), newLog.ActionType, receivedLog.ActionType)
	assert.Equal(s.T(), newLog.Message, receivedLog.Message)
	assert.Equal(s.T(), newLog.Sender, receivedLog.Sender)
	assert.Equal(s.T(), newLog.Status, receivedLog.Status)
	assert.Equal(s.T(), newLog.Language, receivedLog.Language)
	assert.Equal(s.T(), newLog.FullResponse, receivedLog.FullResponse)
	assert.Equal(s.T(), newLog.Created.Unix(), receivedLog.Created.Unix())
	assert.Equal(s.T(), newLog.Created.Unix(), receivedLog.Created.Unix())
	assert.Equal(s.T(), newLog.MessageID, receivedLog.MessageID)
	assert.Equal(s.T(), newLog.StatusDelive, receivedLog.StatusDelive)
	assert.Equal(s.T(), newLog.Cost, receivedLog.Cost)
}

func (s *StoreSuite) TestUpdateLog() {
	newLog := model.CreateTestLog(s.T())
	err := s.Store.LogRepository.InsertSMSLog(newLog)

	logForUpdate := model.CreateLogForUpdate(s.T())
	newLog.UpdateExistLog(&logForUpdate)
	err = s.Store.LogRepository.UpdateSMSLog(newLog)
	s.NoError(err)

	updatedLog, err := s.Store.LogRepository.GetSMSLog(&newLog)
	s.NoError(err)
	assert.Equal(s.T(), logForUpdate.Phone, updatedLog.Phone)
	assert.Equal(s.T(), logForUpdate.Sender, updatedLog.Sender)
	assert.Equal(s.T(), logForUpdate.Status, updatedLog.Status)
	assert.Equal(s.T(), logForUpdate.StatusDelive, updatedLog.StatusDelive)
	assert.Equal(s.T(), logForUpdate.Cost, updatedLog.Cost)
	assert.Equal(s.T(), logForUpdate.Updated.Unix(), updatedLog.Updated.Unix())
	assert.Equal(s.T(), logForUpdate.MessageID, updatedLog.MessageID)
}
