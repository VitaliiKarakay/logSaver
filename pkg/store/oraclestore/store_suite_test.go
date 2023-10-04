package oraclestore_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"logSaver/pkg/config"
	"logSaver/pkg/store/oraclestore"
)

type StoreSuite struct {
	suite.Suite
	Store  *oraclestore.DB
	Oracle *sql.DB

	tables []string
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}

func (s *StoreSuite) SetupSuite() {
	s.tables = []string{
		config.LogTest,
	}
	cfg := config.Config{}
	err := cfg.ReadDatabaseConfigs()
	if err != nil {
		fmt.Println("ReadConfig ", err)
	}

	db, err := oraclestore.New(&cfg)
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
		_, err := s.Store.DB.Exec(`TRUNCATE TABLE ` + table)
		if err != nil {
			fmt.Println("cleanDB ", err)
		}
	}
}
