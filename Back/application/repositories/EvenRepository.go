package repositories

import (
	"github.com/franBarrientos/domain"
)

type IEventRepository interface {
	GetAllEvents() ([]domain.Event, error)

	CreateEvent(e *domain.Event) (domain.Event, error)

	UpdateEvent(id int, e map[string]interface{}) (domain.Event, error)

	GetEventById(id int) (domain.Event, error)

	GetEventsFiltered(date string, state string, title string) ([]domain.Event, error)

	AddSubscribe(subscribe int, event int) error
}
