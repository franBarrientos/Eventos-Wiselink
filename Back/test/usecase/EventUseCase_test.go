package test

import (
	"github.com/franBarrientos/application/usecases"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type EventRepositoryMock struct {
	mock.Mock
}

func (m *EventRepositoryMock) CreateEvent(e *domain.Event) (domain.Event, error) {
	args := m.Called(e)
	return args.Get(0).(domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) UpdateEvent(id int, e map[string]interface{}) (domain.Event, error) {
	args := m.Called(id, e)
	return args.Get(0).(domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) GetEventsFiltered(date string, state string, title string, page int, limit int) ([]domain.Event, error) {
	args := m.Called(date, state, title, page, limit)
	return args.Get(0).([]domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) AddSubscribe(subscribe int, event int) error {
	args := m.Called(subscribe, event)
	return args.Error(0)
}

func (m *EventRepositoryMock) GetAllEvents(page int, limit int) ([]domain.Event, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) GetEventById(id int) (domain.Event, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) GetSubscribersToEvent(eventId int, page int, limit int) ([]domain.User, error) {
	args := m.Called(eventId, page, limit)
	return args.Get(0).([]domain.User), args.Error(1)
}

func TestEventUseCase(t *testing.T) {

	eventRepositoryMock := &EventRepositoryMock{}

	underTest := usecases.NewEventUseCase(eventRepositoryMock)

	t.Run("success UpdateEvent", func(t *testing.T) {

		eventRepositoryMock.On("UpdateEvent", mock.AnythingOfType("int"), mock.IsType(map[string]interface{}{})).Return(domain.Event{
			Title: "update",
		}, nil)

		eventUpdated, err := underTest.UpdateEvent(1, &map[string]interface{}{
			"Title": "update",
		})
		assert.NoError(t, err)
		assert.NotNil(t, eventUpdated)
		assert.Equal(t, "update", eventUpdated.Title)
	})

	t.Run("success GetAllEvents", func(t *testing.T) {
		eventRepositoryMock.On("GetAllEvents", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]domain.Event{}, nil)
		events, errEv := underTest.GetAllEvents(0, 12)

		assert.NoError(t, errEv)
		assert.NotNil(t, events)

	})

	t.Run("success GetEventsFiltered", func(t *testing.T) {
		eventRepositoryMock.On("GetEventsFiltered", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]domain.Event{}, nil)
		eventsFiltered, errFiltered := underTest.GetEventsFiltered("", "", "", 1, 12)
		assert.NoError(t, errFiltered)
		assert.NotNil(t, eventsFiltered)

	})
	t.Run("success CreateEvent", func(t *testing.T) {
		eventRepositoryMock.On("CreateEvent", mock.AnythingOfType("*domain.Event")).Return(domain.Event{}, nil)

		eventCreated, errEc := underTest.CreateEvent(&input.EventAddDTO{
			Date:      &time.Time{},
			Organizer: &input.OrganizerAddDTO{},
			Place:     &input.PlaceAddDTO{},
			State:     new(bool),
		})

		assert.NoError(t, errEc)
		assert.NotNil(t, eventCreated)

	})
	t.Run("SubscribeToEvent", func(t *testing.T) {
		eventRepositoryMock.On("AddSubscribe", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
		errSb := underTest.SubscribeToEvent(&input.SubscribeAddDTO{
			Event: 1,
			User:  1,
		})
		assert.NoError(t, errSb)
		assert.Nil(t, errSb)
	})

}
