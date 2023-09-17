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
	s.setupHTTPServer()

	logData := createTestLog()
	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	response := s.sendHTTPRequest("POST", "/log", requestBody)

	s.assertStatusCode(http.StatusOK, response)
}

func (s *HttpSuite) TestUpdateLog() {
	s.setupHTTPServer()

	logData := createTestLog()
	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	response := s.sendHTTPRequest("POST", "/log", requestBody)

	s.assertStatusCode(http.StatusOK, response)

	updateExistLog(logData)

	response = s.sendHTTPRequest("PUT", "/log", requestBody)

	s.assertStatusCode(http.StatusOK, response)
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

func (s *HttpSuite) sendHTTPRequest(method, path string, body []byte) int {
	url := "http://localhost:8080" + path
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	s.NoError(err)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		s.NoError(err)
	}(resp.Body)

	return resp.StatusCode
}

func (s *HttpSuite) assertStatusCode(expected int, responseCode int) {
	assert.Equal(s.T(), expected, responseCode)
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
