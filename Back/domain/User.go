package domain

import (
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
	Events    []Event
}

type IUserUseCase interface {
	RegisterUser(user *input.UserAddDTO) (output.AuthResponse, error)
	LoginUser(email, password string) (output.AuthResponse, error)
	GetEventsSubscribed(idUser int, state string) ([]output.EventDTO, error)
}
