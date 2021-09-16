package pkg

import (
	"time"
)

var Events []Event
var Id int

type Event struct {
	EventId     int    `json:"event_id"`
	UserId      int    `json:"user_id"`
	Date        MyTime `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MyTime struct {
	time.Time
}

func (m *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` || string(data) == "" {
		*m = MyTime{time.Now()}
		return nil
	}
	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = MyTime{tt}
	return err
}
