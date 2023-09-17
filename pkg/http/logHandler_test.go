package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (s *HttpSuite) TestCreateLog() {
	r := gin.Default()
	r.POST("/log", s.logHandler.CreateLog)

	go func() {
		if err := r.Run(":8080"); err != nil {
			s.NoError(err)
		}
	}()

	logData := createTestLog()

	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	req, err := http.NewRequest("POST", "http://localhost:8080/log", bytes.NewReader(requestBody))
	s.NoError(err)

	resp, err := http.DefaultClient.Do(req)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	s.NoError(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		s.NoError(err)
	}(resp.Body)
}

func (s *HttpSuite) TestUpdateLog() {
	r := gin.Default()
	r.POST("/log", s.logHandler.CreateLog)
	r.PUT("/log", s.logHandler.UpdateLog)

	go func() {
		if err := r.Run(":8080"); err != nil {
			s.NoError(err)
		}
	}()

	logData := createTestLog()

	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	req, err := http.NewRequest("POST", "http://localhost:8080/log", bytes.NewReader(requestBody))
	s.NoError(err)

	resp, err := http.DefaultClient.Do(req)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	s.NoError(err)

	updateExistLog(logData)

	requestBody, err = json.Marshal(logData)
	s.NoError(err)

	req, err = http.NewRequest("PUT", "http://localhost:8080/log", bytes.NewReader(requestBody))
	s.NoError(err)

	resp, err = http.DefaultClient.Do(req)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	s.NoError(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		s.NoError(err)
	}(resp.Body)
}

func createTestLog() map[string]interface{} {
	logData := map[string]interface{}{
		"userID":       2134496917,
		"phone":        "380953071221",
		"actionID":     324,
		"actionTitle":  "Good Action",
		"actionType":   "promoSMS",
		"message":      "some message",
		"sender":       "intistele",
		"status":       "success",
		"language":     "en",
		"fullResponse": "sms_id: 6774560000068401360004",
		"created":      "2023-02-27T00:27:00.031Z",
		"updated":      "2023-02-27T00:27:00.031Z",
		"messageID":    "6774560000068401360005",
		"StatusDelive": 1,
		"cost":         14.88,
	}
	return logData
}

func updateExistLog(logData map[string]interface{}) {
	logData["StatusDelive"] = 2
	logData["Cost"] = 3.22
}
