package repositories

import (
	"github.com/franBarrientos/domain"
)

type IUserRepository interface {
	CreateUser(user *domain.User) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	GetUserById(id int) (domain.User, error)
	GetEventsSubscribed(idUser int, state string, page int, limit int) ([]domain.Event, error)
}
