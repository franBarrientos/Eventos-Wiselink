package mappers_db

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/infrastructure/gorm/entities_db"
)

func UserDomainToUserEntity(user *domain.User) entities_db.User {
	return entities_db.User{
		Email:    user.Email,
		Password: user.Password,
		Role:     domain.USER,
		PersonalData: entities_db.PersonalData{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
}

func UserEntityToUserDomain(user *entities_db.User) domain.User {

	events := []domain.Event{}
	for _, event := range user.EventsSubscribed {
		events = append(events, EventEntityToEventDomain(&event))
	}

	return domain.User{
		Id:        user.ID,
		FirstName: user.PersonalData.FirstName,
		LastName:  user.PersonalData.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      string(user.Role),
		Events:    events,
	}

}
