package domain

import (
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
	GetEventsSubscribed(idUser int, state string, page int, limit int) ([]output.EventDTO, error)
}
