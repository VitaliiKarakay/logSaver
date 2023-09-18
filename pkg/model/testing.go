// nolint
package model

import (
	"fmt"
	"testing"
	"time"
)

func CreateTestLog(t *testing.T) Log {
	t.Helper()
	created, err := time.Parse(time.RFC3339, "2023-02-27T00:27:00.031+02:00")
	if err != nil {
		fmt.Println("Parsing created data error: ", err)
	}
	updated, err := time.Parse(time.RFC3339, "2023-02-27T00:27:00.031+02:00")
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

func CreateLogForUpdate(t *testing.T) Log {
	t.Helper()

	updated, err := time.Parse(time.RFC3339, "22023-02-27T06:27:05.097Z")
	if err != nil {
		fmt.Println("Parsing updated data error: ", err)
	}

	return Log{
		Phone:        "380953071221",
		Sender:       "intistele",
		Status:       "success",
		StatusDelive: 2,
		Cost:         3.8351,
		Updated:      updated,
		MessageID:    "6774560000068401360008",
	}
}
