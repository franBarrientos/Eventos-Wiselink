package repositories_db

import (
	"github.com/franBarrientos/application/repositories"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/infrastructure/gorm/entities_db"
	"github.com/franBarrientos/infrastructure/gorm/mappers_db"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryDb struct {
	database *gorm.DB
}

func NewUserRepositoryDb(db *gorm.DB) repositories.IUserRepository {
	return &UserRepositoryDb{
		database: db,
	}
}

func (u UserRepositoryDb) CreateUser(user *domain.User) (domain.User, error) {
	userToCreate := mappers_db.UserDomainToUserEntity(user)
	result := u.database.Create(&userToCreate)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return mappers_db.UserEntityToUserDomain(&userToCreate), nil

}

func (u UserRepositoryDb) GetUserByEmail(email string) (domain.User, error) {
	var user entities_db.User
	result := u.database.Where("email = ?", email).Preload("PersonalData").Preload("EventsSubscribed").Preload("EventsSubscribed.Place").Preload("EventsSubscribed.Organizer").Preload("EventsSubscribed.Organizer.PersonalData").First(&user)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return mappers_db.UserEntityToUserDomain(&user), nil

}

func (u UserRepositoryDb) GetUserById(id int) (domain.User, error) {
	var user entities_db.User
	result := u.database.Preload("EventsSubscribed").Preload("EventsSubscribed.Place").Preload("EventsSubscribed.Organizer").Preload("EventsSubscribed.Organizer.PersonalData").Preload("PersonalData").Where("id = ?", id).First(&user)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return mappers_db.UserEntityToUserDomain(&user), nil

}

func (u UserRepositoryDb) GetEventsSubscribed(idUser int, state string, page int, limit int) ([]domain.Event, error) {
	var events []domain.Event
	var user entities_db.User
	result := u.database.Preload("EventsSubscribed").Preload("EventsSubscribed.Place").Preload("EventsSubscribed.Organizer").Preload("EventsSubscribed.Organizer.PersonalData").Preload("PersonalData").Where("id = ?", idUser).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, e := range user.EventsSubscribed {
		if state == "active" && e.Date.After(time.Now()) {
			events = append(events, mappers_db.EventEntityToEventDomain(&e))
		}
		if state == "completed" && e.Date.Before(time.Now()) {
			events = append(events, mappers_db.EventEntityToEventDomain(&e))
		}
		if state == "" {
			events = append(events, mappers_db.EventEntityToEventDomain(&e))
		}
	}

	return events, nil
}
