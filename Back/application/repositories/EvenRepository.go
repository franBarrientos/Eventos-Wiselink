package repositories

import (
	"github.com/franBarrientos/domain"
)

type IEventRepository interface {
	GetAllEvents(page int, limit int) ([]domain.Event, error)

	CreateEvent(e *domain.Event) (domain.Event, error)

	UpdateEvent(id int, e map[string]interface{}) (domain.Event, error)

	GetEventById(id int) (domain.Event, error)

	GetEventsFiltered(date string, state string, title string, page int, limit int) ([]domain.Event, error)

	AddSubscribe(subscribe int, event int) error

	GetSubscribersToEvent(eventId int, page int, limit int) ([]domain.User, error)
}
