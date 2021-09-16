package services

import (
	"11/pkg"
	"fmt"
	"reflect"
	"time"
)

type EventsService struct {
}

func NewEventsService() *EventsService {
	return &EventsService{}
}

func (e *EventsService) CreateEvent(eventId int, userId int, date pkg.MyTime, title, description string) error {
	if reflect.DeepEqual(date.Time, time.Time{}) {
		date.Time = time.Now()
	}
	eventId = pkg.Id
	pkg.Id++
	event := pkg.Event{
		EventId:     eventId,
		UserId:      userId,
		Date:        date,
		Title:       title,
		Description: description,
	}
	pkg.Events = append(pkg.Events, event)
	return nil
}

func (e *EventsService) UpdateEvent(eventId int, userId int, date pkg.MyTime, title, description string) error {
	var valid bool // проверяем существует ли событие
	var index int  // запоминаем индекс, если существует
	for i, v := range pkg.Events {
		if v.EventId == eventId {
			valid = true
			index = i
			break
		}
	}
	if !valid {
		return fmt.Errorf("event doen't exist")
	}
	if reflect.DeepEqual(date.Time, time.Time{}) {
		date.Time = time.Now()
	}
	pkg.Events[index].UserId = userId
	pkg.Events[index].Date = date
	pkg.Events[index].Title = title
	pkg.Events[index].Description = description
	return nil
}

func (e *EventsService) DeleteEvent(eventId int) error {
	var valid bool // проверяем существует ли событие
	var index int  // запоминаем индекс, если существует
	for i, v := range pkg.Events {
		if v.EventId == eventId {
			valid = true
			index = i
			break
		}
	}
	if !valid {
		return fmt.Errorf("event doen't exist")
	}
	for i := index; i < len(pkg.Events)-1; i++ {
		pkg.Events[i].EventId = pkg.Events[i+1].EventId - 1
		pkg.Events[i].UserId = pkg.Events[i+1].UserId
		pkg.Events[i].Date = pkg.Events[i+1].Date
		pkg.Events[i].Title = pkg.Events[i+1].Title
		pkg.Events[i].Description = pkg.Events[i+1].Description
	}
	pkg.Events = pkg.Events[:len(pkg.Events)-1]
	pkg.Id--
	return nil
}

func (e *EventsService) EventsForDay(date pkg.MyTime) ([]pkg.Event, error) {
	if reflect.DeepEqual(date.Time, time.Time{}) {
		date.Time = time.Now()
	}
	var res []pkg.Event
	for _, v := range pkg.Events {
		if date.Day() == v.Date.Day() && date.Month() == v.Date.Month() && date.Year() == v.Date.Year() {
			res = append(res, v)
		}
	}
	return res, nil
}

func (e *EventsService) EventsForWeek(date pkg.MyTime) ([]pkg.Event, error) {
	if reflect.DeepEqual(date.Time, time.Time{}) {
		date.Time = time.Now()
	}
	var res []pkg.Event
	for _, v := range pkg.Events {
		y1, w1 := date.ISOWeek()
		y2, w2 := v.Date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			res = append(res, v)
		}
	}
	return res, nil
}

func (e *EventsService) EventsForMonth(date pkg.MyTime) ([]pkg.Event, error) {
	if reflect.DeepEqual(date.Time, time.Time{}) {
		date.Time = time.Now()
	}
	var res []pkg.Event
	for _, v := range pkg.Events {
		if date.Month() == v.Date.Month() && date.Year() == v.Date.Year() {
			res = append(res, v)
		}
	}
	return res, nil
}
