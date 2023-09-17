package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (s *HttpSuite) TestCreateLog() {
	//s.setupHTTPServer()

	logData := createTestLog()
	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	response := s.sendHTTPRequest("POST", "/log", requestBody)

	s.assertStatusCode(http.StatusOK, response)
}

func (s *HttpSuite) TestUpdateLog() {
	//s.setupHTTPServer()

	logData := createTestLog()
	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	response := s.sendHTTPRequest("POST", "/log", requestBody)

	s.assertStatusCode(http.StatusOK, response)

	updateExistLog(logData)
	requestBody, err = json.Marshal(logData)
	s.NoError(err)

	response = s.sendHTTPRequest("PUT", "/log", requestBody)

	s.assertStatusCode(http.StatusOK, response)
}

func (s *HttpSuite) TestUpdateIncorrectLog() {
	//s.setupHTTPServer()

	logData := createTestLog()
	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	response := s.sendHTTPRequest("POST", "/log", requestBody)

	s.assertStatusCode(http.StatusOK, response)

	updateExistLog(logData)
	logData["MessageID"] = "6774560000068401360001"
	requestBody, err = json.Marshal(logData)
	s.NoError(err)

	response = s.sendHTTPRequest("PUT", "/log", requestBody)

	s.assertStatusCode(http.StatusInternalServerError, response)
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
		"UserID":       2134496917,
		"Phone":        "380953071221",
		"ActionID":     324,
		"ActionTitle":  "Good Action",
		"ActionType":   "promoSMS",
		"Message":      "some message",
		"Sender":       "intistele",
		"Status":       "success",
		"Language":     "en",
		"FullResponse": "sms_id: 6774560000068401360004",
		"Created":      "2023-02-27T00:27:00.031Z",
		"Updated":      "2023-02-27T00:27:00.031Z",
		"MessageID":    "6774560000068401360005",
		"StatusDelive": 1,
		"Cost":         14.88,
	}

	return logData
}

func updateExistLog(logData map[string]interface{}) {
	logData["StatusDelive"] = 2
	logData["Cost"] = 3.22
}
