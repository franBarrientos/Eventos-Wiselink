package controllers

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
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) LoginUser(credential *input.LoginDTO) (output.AuthResponse, error) {
	args := m.Called(credential)
	return args.Get(0).(output.AuthResponse), args.Error(1)
}
func (m *AuthUseCaseMock) RegisterUser(user *input.UserAddDTO) (output.AuthResponse, error) {
	args := m.Called(user)
	return args.Get(0).(output.AuthResponse), args.Error(1)
}

func TestAuthController(t *testing.T) {

	//mock
	mockAuthUseCase := new(AuthUseCaseMock)
	mockAuthUseCase.On("LoginUser", &input.LoginDTO{
		Email:    "wrong@gmail.com",
		Password: "wrong",
	}).Return(output.AuthResponse{}, errors.New("invalid credentials"))

	mockAuthUseCase.On("LoginUser", &input.LoginDTO{
		Email:    "success@gmail.com",
		Password: "success",
	}).Return(output.AuthResponse{
		User: output.UserDTO{
			FirstName: "string",
			LastName:  "string",
			Email:     "string",
			Role:      "USER",
		},
		Token: output.LoginResponse{
			AccessToken:  "string",
			RefreshToken: "string",
		},
	}, nil)

	authController := controllers.AuthController{
		AuthUseCase: mockAuthUseCase,
		Validator:   validator.New(validator.WithRequiredStructEnabled()),
	}

	t.Run("Success Login", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)
		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(`{"Email": "success@gmail.com", "Password": "success"}`)
		err := authController.Login(c)
		assert.NoError(t, err)
	})

	t.Run("Wrong Login", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)
		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(`{"Email": "wrong@gmail.com", "Password": "wrong"}`)
		authController.Login(c)
		assert.Equal(t, fiber.StatusInternalServerError, c.Response().StatusCode())
	})

	t.Run("Wrong Email", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)
		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(`{"Email": "wronmail.com", "Password": "wrong"}`)
		authController.Login(c)
		assert.Equal(t, fiber.StatusBadRequest, c.Response().StatusCode())
	})
	t.Run("Empty Email", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)
		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(`{"Email": "", "Password": "wrong"}`)
		authController.Login(c)
		assert.Equal(t, fiber.StatusBadRequest, c.Response().StatusCode())
	})
	t.Run("Empty Password", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)
		c.Request().Header.SetContentType("application/json")
		c.Request().SetBodyString(`{"Email": "wronmail.com", "Password": ""}`)
		authController.Login(c)
		assert.Equal(t, fiber.StatusBadRequest, c.Response().StatusCode())
	})

}
