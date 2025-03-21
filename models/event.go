package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"` //chỗ này nếu mình không điền thì báo lỗi
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	// later: add ot tp a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
