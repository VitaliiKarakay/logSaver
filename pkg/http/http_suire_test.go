package http_test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"logSaver/pkg/config"
	"logSaver/pkg/http"
	_ "logSaver/pkg/store"
	"logSaver/pkg/store/mongostore"
	_ "logSaver/pkg/store/postgresstore"
)

type HttpSuite struct {
	suite.Suite
	Store      *mongostore.DB
	logHandler http.LogHandler

	tables []string
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(HttpSuite))
}

func (s *HttpSuite) SetupSuite() {
	s.tables = []string{
		config.LogTest,
	}

	cfg, err := config.ReadMongoConfig()
	if err != nil {
		fmt.Println("ReadConfig ", err)
	}

	db, err := mongostore.New(cfg)
	if err != nil {
		fmt.Println("store.New ", err)
	}

	s.logHandler = http.LogHandler{DB: db}

	s.Store = db
	s.setupHTTPServer()
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
	lr := mongostore.NewLogRepository(s.Store.DB)

	err := lr.DeleteAllLogs()
	if err != nil {
		// Обработайте ошибку, если необходимо
		fmt.Println("Ошибка при очистке коллекции: ", err)
	}
}

func (s *HttpSuite) setupHTTPServer() {
	r := gin.Default()
	r.POST("/log", s.logHandler.CreateLog)
	r.PUT("/log", s.logHandler.UpdateLog)

	go func() {
		if err := r.Run(":8080"); err != nil {
			s.NoError(err)
		}
	}()
}
