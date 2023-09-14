package store_test

import (
	"logSaver/pkg/model"
)

func (s *StoreSuite) TestLogInsert() {
	log := model.GetTestLog(s.T())

	err := s.Store.LogRepository.Insert(log)
	s.NoError(err)
}
