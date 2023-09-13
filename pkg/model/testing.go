package model

import (
	"fmt"
	"testing"
	"time"
)

func GetTestLog(t *testing.T) Log {
	t.Helper()
	created, err := time.Parse(time.RFC3339, "2023-02-27T00:27:00.031Z")
	if err != nil {
		fmt.Println("Parsing created data error: ", err)
	}
	updated, err := time.Parse(time.RFC3339, "2023-02-27T00:27:00.031Z")
	if err != nil {
		fmt.Println("Parsing updated data error: ", err)
	}

	return Log{
		UserID:       2134496917,
		Phone:        "380953071221",
		ActionID:     324,
		ActionTitle:  "Good Action",
		ActionType:   "promoSMS",
		Message:      "some message",
		Sender:       "intistele",
		Status:       "success",
		Language:     "en",
		FullResponse: "sms_id: 6774560000068401360004",
		Created:      created,
		Updated:      updated,
		MessageID:    "6774560000068401360008",
		StatusDelive: 1,
		Cost:         2.28,
	}
}
