package domain

import (
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
	"time"
)

type Event struct {
	Id               int
	Title            string
	ShortDescription string
	LongDescription  string
	Date             time.Time
	Organizer        Organizer
	Place            Place
	State            bool
}

type IEventUseCase interface {
	GetAllEvents() ([]output.EventDTO, error)

	CreateEvent(event *input.EventAddDTO) (output.EventDTO, error)

	UpdateEvent(id int, event *map[string]interface{}) (output.EventDTO, error)

	GetEventsFiltered(date string, state string, title string) ([]output.EventDTO, error)

	SubscribeToEvent(subscribe *input.SubscribeAddDTO) error
}
