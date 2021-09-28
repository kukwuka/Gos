package services

import (
	"github.com/kukwuka/Gos/develop/3_11/pkg"
)

type IEvent interface {
	CreateEvent(eventId int, userId int, date pkg.MyTime, title, description string) error
	UpdateEvent(eventId int, userId int, date pkg.MyTime, title, description string) error
	DeleteEvent(eventId int) error
	EventsForDay(date pkg.MyTime) ([]pkg.Event, error)
	EventsForWeek(date pkg.MyTime) ([]pkg.Event, error)
	EventsForMonth(date pkg.MyTime) ([]pkg.Event, error)
}

type Service struct {
	IEvent
}

func NewService() *Service {
	return &Service{
		NewEventsService(),
	}
}
