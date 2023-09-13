package store_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"logSaver/pkg/config"
	"logSaver/pkg/store"
)

type StoreSuite struct {
	suite.Suite
	Store *store.DB

	tables []string
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}

func (s *StoreSuite) SetupSuite() {
	s.tables = []string{
		"logTest", //не забыть вынести в константу
	}
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("ReadConfig ", err)
	}

	db, err := store.New(cfg)
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
	for _, table := range s.tables {
		_, err := s.Store.Oracle.Exec(`TRUNCATE TABLE ` + table)
		if err != nil {
			fmt.Println("cleanDB ", err)
		}
	}
}
