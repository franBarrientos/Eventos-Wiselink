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
	args := m.Called()
	return args.Get(0).(domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) UpdateEvent(id int, e map[string]interface{}) (domain.Event, error) {
	args := m.Called()
	return args.Get(0).(domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) GetEventsFiltered(date string, state string, title string, page int, limit int) ([]domain.Event, error) {
	args := m.Called()
	return args.Get(0).([]domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) AddSubscribe(subscribe int, event int) error {
	args := m.Called()
	return args.Error(0)
}

func (m *EventRepositoryMock) GetAllEvents(page int, limit int) ([]domain.Event, error) {
	args := m.Called()
	return args.Get(0).([]domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) GetEventById(id int) (domain.Event, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) GetSubscribersToEvent(eventId int, page int, limit int) ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func TestEventUseCase(t *testing.T) {

	mockEventList := []domain.Event{
		{
			Id:               1,
			Title:            "fsdfs",
			ShortDescription: "fsdfs",
			LongDescription:  "fsdfs",
			Date:             time.Now(),
			Organizer: domain.Organizer{
				Id:        20,
				FirstName: "fs",
				LastName:  "fs",
			},
			Place: domain.Place{
				Id:            2,
				Address:       "fs",
				AddressNumber: 222,
				City:          "fs",
				Country:       "fs",
			},
			State: false,
		},
	}

	t.Run("success", func(t *testing.T) {

		eventRepositoryMock := &EventRepositoryMock{}

		eventRepositoryMock.On("GetAllEvents").Return(mockEventList, nil)

		eventRepositoryMock.On("GetEventById", mock.AnythingOfType("int")).Return(domain.Event{}, nil)

		eventRepositoryMock.On("CreateEvent").Return(mockEventList[0], nil)

		eventRepositoryMock.On("UpdateEvent").Return(domain.Event{}, nil)

		eventRepositoryMock.On("GetEventsFiltered").Return(mockEventList, nil)

		eventRepositoryMock.On("AddSubscribe").Return(nil)

		//mock del repositorio
		underTest := usecases.NewEventUseCase(eventRepositoryMock)

		eventUpdated, err := underTest.UpdateEvent(1, &map[string]interface{}{
			"Title": "update",
		})

		assert.NoError(t, err)
		assert.NotNil(t, eventUpdated)

		events, errEv := underTest.GetAllEvents(0, 12)

		assert.NoError(t, errEv)
		assert.NotNil(t, events)

		eventsFiltered, errFiltered := underTest.GetEventsFiltered("", "", "", 1, 12)
		assert.NoError(t, errFiltered)
		assert.NotNil(t, eventsFiltered)

		vf := false
		eventCreated, errEc := underTest.CreateEvent(&input.EventAddDTO{
			Title:            "fsdfs",
			ShortDescription: "fsdfs",
			LongDescription:  "fsdfs",
			Date:             &time.Time{},
			Organizer: &input.OrganizerAddDTO{
				FirstName: "fsd",
				LastName:  "fsd",
			},
			Place: &input.PlaceAddDTO{
				Address:       "fsd",
				AddressNumber: 0,
				City:          "fds",
				Country:       "fds",
			},
			State: &vf,
		})

		assert.NoError(t, errEc)
		assert.NotNil(t, eventCreated)

		errSb := underTest.SubscribeToEvent(&input.SubscribeAddDTO{
			Event: 1,
			User:  1,
		})

		assert.NoError(t, errSb)

	})

}
