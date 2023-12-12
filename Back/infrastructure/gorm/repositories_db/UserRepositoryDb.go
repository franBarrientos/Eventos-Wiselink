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
	result := u.database.Where("email = ?", email).
		Preload("PersonalData").
		Preload("EventsSubscribed").
		Preload("EventsSubscribed.Place").
		Preload("EventsSubscribed.Organizer").
		Preload("EventsSubscribed.Organizer.PersonalData").
		First(&user)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return mappers_db.UserEntityToUserDomain(&user), nil

}

func (u UserRepositoryDb) GetUserById(id int) (domain.User, error) {
	var user entities_db.User
	result := u.database.Preload("EventsSubscribed").
		Preload("EventsSubscribed.Place").
		Preload("EventsSubscribed.Organizer").
		Preload("EventsSubscribed.Organizer.PersonalData").
		Preload("PersonalData").
		Where("id = ?", id).
		First(&user)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return mappers_db.UserEntityToUserDomain(&user), nil

}

func (u UserRepositoryDb) GetEventsSubscribed(idUser int, state string, page int, limit int) ([]domain.Event, error) {
	events := []entities_db.Event{}

	query := u.database.Model(&entities_db.User{}).
		Select("events.*, users.id as user_id").
		Joins("JOIN user_events ON users.id = user_events.user_id").
		Joins("JOIN events ON user_events.event_id = events.id").
		Where("users.id = ?", idUser).
		Preload("Place").
		Preload("Organizer").
		Preload("Organizer.PersonalData").
		Offset((page - 1) * limit).
		Limit(limit)

	switch state {
	case "active":
		query = query.Where("events.date > ?", time.Now())
	case "completed":
		query = query.Where("events.date <= ?", time.Now())
	}

	result := query.Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}

	eventsDomain := []domain.Event{}
	for _, i := range events {
		eventsDomain = append(eventsDomain, mappers_db.EventEntityToEventDomain(&i))
	}

	return eventsDomain, nil
}
