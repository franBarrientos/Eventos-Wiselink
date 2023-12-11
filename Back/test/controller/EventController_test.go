package controllers_test

import (
	"errors"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
	"github.com/franBarrientos/infrastructure/rest/controllers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
	"testing"
	"time"
)

// MockEventUseCase mocks the IEventUseCase interface
type MockEventUseCase struct {
	mock.Mock
}

func (m *MockEventUseCase) GetAllEvents(page int, limit int) ([]output.EventDTO, error) {
	args := m.Called()
	return args.Get(0).([]output.EventDTO), args.Error(1)
}

func (m *MockEventUseCase) CreateEvent(event *input.EventAddDTO) (output.EventDTO, error) {
	args := m.Called(event)
	return args.Get(0).(output.EventDTO), args.Error(1)
}

func (m *MockEventUseCase) GetEventsFiltered(date string, state string, title string) ([]output.EventDTO, error) {
	args := m.Called(date, state, title)
	return args.Get(0).([]output.EventDTO), args.Error(1)
}

func (m *MockEventUseCase) UpdateEvent(id int, event *map[string]interface{}) (output.EventDTO, error) {
	args := m.Called(id, event)
	return args.Get(0).(output.EventDTO), args.Error(1)
}

func (m *MockEventUseCase) SubscribeToEvent(subscribe *input.SubscribeAddDTO) error {
	args := m.Called(subscribe)
	return args.Error(0)
}

// MockUserUseCase mocks the IUserUseCase interface
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) GetEventsSubscribed(idUser int, state string) ([]output.EventDTO, error) {
	args := m.Called(idUser, state)
	return args.Get(0).([]output.EventDTO), args.Error(1)
}

// MockValidator mocks the validator.Validate interface
type MockValidator struct {
	mock.Mock
}

func (m *MockValidator) Struct(s interface{}) error {
	args := m.Called(s)
	return args.Error(0)
}

func TestEventController_Methods(t *testing.T) {
	eventUseCase := new(MockEventUseCase)
	userUseCase := new(MockUserUseCase)
	validator := validator.New(validator.WithRequiredStructEnabled())

	controller := controllers.EventController{
		EventUseCase: eventUseCase,
		UserUseCase:  userUseCase,
		Validator:    validator,
	}

	t.Run("GetAllEvents", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		eventUseCase.On("GetAllEvents").Return([]output.EventDTO{}, nil)

		err := controller.GetAllEvents(c) // Assuming `c` is a valid Fiber context

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, c.Response().StatusCode())
	})

	t.Run("CreateEvent Success", func(t *testing.T) {

		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		// Mock data for the request
		requestBody := ` {
        "Title": "aaaa2",
        "ShortDescription": "loooooo2",
        "LongDescription": "laaaa",
        "Date": "2023-12-06T14:18:11-03:00",
        "Organizer": {
            "FirstName": "camila",
            "LastName": "barr"
        },
        "Place": {
            "Address": "jamaica 2",
            "AddressNumber": 4207,
            "City": "corrientes2",
            "Country": "argentina2"
        },
        "State": true
    	}`

		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(requestBody)

		eventUseCase.On("CreateEvent", mock.AnythingOfType("*input.EventAddDTO")).Return(output.EventDTO{
			Id:               1,
			Title:            "aaaa",
			ShortDescription: "loooooo",
			LongDescription:  "laaaa",
			Date:             time.Now(),
			Organizer: output.OrganizerDTO{
				Id:        1,
				FirstName: "camila",
				LastName:  "barr",
			},
			Place: output.PlaceDTO{
				Id:            1,
				Address:       "jamaica 2",
				AddressNumber: 4207,
				City:          "corrientes2",
				Country:       "argentina2",
			},
			State: true,
		}, nil)

		errorCreared := controller.CreateEvent(c)

		assert.NoError(t, errorCreared)
		assert.Equal(t, fiber.StatusCreated, c.Response().StatusCode())

	})
	t.Run("CreateEvent Wrong", func(t *testing.T) {

		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		requestBody := ` {
        "Title": "",
        "ShortDescription": "",
        "LongDescription": "",
        "Date": null,
        "Organizer": null,
        "Place": null,
        "State": false
    	}`

		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(requestBody)

		eventUseCase.On("CreateEvent", &input.EventAddDTO{
			Title:            "",
			ShortDescription: "",
			LongDescription:  "",
			Date:             nil,
			Organizer:        nil,
			Place:            nil,
			State:            new(bool),
		}).Return(output.EventDTO{}, errors.New(" error"))

		controller.CreateEvent(c)

		assert.Equal(t, fiber.StatusBadRequest, c.Response().StatusCode())

	})

	t.Run("Subscribe Success", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		requestBody := ` {
        "User": 2,
        "Event": 3
    	}`

		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(requestBody)

		eventUseCase.On("SubscribeToEvent", &input.SubscribeAddDTO{
			User:  2,
			Event: 3,
		}).Return(nil)

		userUseCase.On("GetEventsSubscribed", 2, "").Return([]output.EventDTO{}, nil)

		controller.SubscribeUserToEvent(c)

		assert.Equal(t, fiber.StatusOK, c.Response().StatusCode())

	})
	t.Run("Subscribe Wrong", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)

		// Mock data for the request
		requestBody := ` {
        "User": 4,
        "Event": 4
    	}`

		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(requestBody)

		eventUseCase.On("SubscribeToEvent", &input.SubscribeAddDTO{
			User:  4,
			Event: 4,
		}).Return(errors.New(" error"))

		controller.SubscribeUserToEvent(c)

		assert.Equal(t, fiber.StatusInternalServerError, c.Response().StatusCode())

	})

}
