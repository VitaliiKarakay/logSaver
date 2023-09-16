package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	http2 "logSaver/pkg/http"
)

func (s *HttpSuite) TestCreateLog() {
	r := gin.Default()
	handler := http2.LogHandler{DB: s.Store} // Здесь создайте экземпляр вашего обработчика
	r.POST("/log", handler.CreateLog)

	go func() {
		if err := r.Run(":8080"); err != nil {
			s.T().Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

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

	requestBody, err := json.Marshal(logData)
	s.NoError(err)

	req, err := http.NewRequest("POST", "http://localhost:8080/log", bytes.NewReader(requestBody))
	s.NoError(err)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		s.NoError(err)
	}(resp.Body)

	s.NoError(err)
}

func (s *HttpSuite) TestUpdateLog() {

}
