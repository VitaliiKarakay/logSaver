package mongostore_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"logSaver/pkg/config"
	"logSaver/pkg/store/mongostore"
)

type StoreSuite struct {
	suite.Suite
	Store *mongostore.DB
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}

func (s *StoreSuite) SetupSuite() {
	cfg, err := config.ReadMongoConfig()
	if err != nil {
		fmt.Println("ReadConfig ", err)
	}

	db, err := mongostore.New(cfg)
	if err != nil {
		fmt.Println("store.New ", err)
	}

	s.Store = db
}

func (s *StoreSuite) BeforeTest() {
	s.cleanDB()
}

func (s *StoreSuite) TearDownTest() {
	s.cleanDB()
}

func (s *StoreSuite) TearDownSuite() {
	s.cleanDB()
}

func (s *StoreSuite) cleanDB() {
	lr := mongostore.NewLogRepository(s.Store.DB)

	err := lr.DeleteAllLogs()
	if err != nil {
		fmt.Println("Error while cleaning the DB ", err)
	}
}
