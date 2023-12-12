package test

import (
	"testing"
	"time"

	"github.com/franBarrientos/application/usecases"
	"github.com/franBarrientos/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetEventsSubscribed(idUser int, state string, page int, limit int) ([]domain.Event, error) {
	args := m.Called(idUser, state, page, limit)
	return args.Get(0).([]domain.Event), args.Error(1)
}

func (m *UserRepositoryMock) CreateUser(user *domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserById(id int) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestUserUseCase_GetEventsSubscribed(t *testing.T) {
	// Create a mock for the UserRepository
	userRepositoryMock := &UserRepositoryMock{}

	// Create a user with subscribed events for testing
	mockUser := domain.User{
		Id: 1,
		Events: []domain.Event{
			{
				Id:        1,
				Title:     "Event 1",
				Date:      time.Now().AddDate(0, 0, 1), // Event in the future
				Organizer: domain.Organizer{},
				Place:     domain.Place{},
				State:     true,
			},
			{
				Id:        2,
				Title:     "Event 2",
				Date:      time.Now().AddDate(0, 0, -1), // Event in the past
				Organizer: domain.Organizer{},
				Place:     domain.Place{},
				State:     true,
			},
		},
	}

	// Configure the mock to return the mockUser when GetUserById is called
	userRepositoryMock.On("GetUserById", mock.AnythingOfType("int")).Return(mockUser, nil)
	userRepositoryMock.On("GetEventsSubscribed", 1, "active", 1, 1).Return([]domain.Event{mockUser.Events[0]}, nil)
	userRepositoryMock.On("GetEventsSubscribed", 1, "completed", 1, 1).Return([]domain.Event{mockUser.Events[1]}, nil)
	userRepositoryMock.On("GetEventsSubscribed", 1, "", 1, 2).Return(mockUser.Events, nil)
	// Create the user use case with the mock repository
	userUseCase := usecases.NewUserUseCase(userRepositoryMock)

	// Test the GetEventsSubscribed method
	t.Run("Get active events", func(t *testing.T) {
		events, err := userUseCase.GetEventsSubscribed(1, "active", 1, 1)
		assert.NoError(t, err)
		assert.Len(t, events, 1) // Only the future event should be returned
		assert.Equal(t, "Event 1", events[0].Title)
	})

	t.Run("Get completed events", func(t *testing.T) {
		events, err := userUseCase.GetEventsSubscribed(1, "completed", 1, 1)
		assert.NoError(t, err)
		assert.Len(t, events, 1) // Only the past event should be returned
		assert.Equal(t, "Event 2", events[0].Title)
	})

	t.Run("Get all events", func(t *testing.T) {
		events, err := userUseCase.GetEventsSubscribed(1, "", 1, 2)
		assert.NoError(t, err)
		assert.Len(t, events, 2) // Both events should be returned
		assert.Equal(t, "Event 1", events[0].Title)
		assert.Equal(t, "Event 2", events[1].Title)
	})
}
