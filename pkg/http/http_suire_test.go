package http_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"logSaver/pkg/config"
	"logSaver/pkg/store"
)

type HttpSuite struct {
	suite.Suite
	Store *store.DB

	tables []string
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(HttpSuite))
}

func (s *HttpSuite) SetupSuite() {
	s.tables = []string{
		config.LogTest,
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

func (s *HttpSuite) BeforeTest() {
	s.cleanDB()
}

func (s *HttpSuite) TearDownTest() {
	s.cleanDB()
}

func (s *HttpSuite) TearDownSuite() {
	s.cleanDB()
}

func (s *HttpSuite) cleanDB() {
	for _, table := range s.tables {
		_, err := s.Store.Oracle.Exec(`TRUNCATE TABLE ` + table)
		if err != nil {
			fmt.Println("cleanDB ", err)
		}
	}
}
