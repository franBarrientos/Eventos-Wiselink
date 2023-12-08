package usecases

import (
	"github.com/franBarrientos/application/mappers_dto"
	"github.com/franBarrientos/application/repositories"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/output"
	"time"
)

type UserUseCase struct {
	userRepository repositories.IUserRepository
}

func NewUserUseCase(userRepository repositories.IUserRepository) domain.IUserUseCase {

	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u UserUseCase) GetEventsSubscribed(idUser int, state string) ([]output.EventDTO, error) {

	result, err := u.userRepository.GetUserById(idUser)
	if err != nil {
		return nil, err
	}
	var events []output.EventDTO
	for _, event := range result.Events {
		if state == "" {
			events = append(events, mappers_dto.EventDomainToEventDTO(&event))
		} else if state == "active" && event.Date.After(time.Now()) {
			events = append(events, mappers_dto.EventDomainToEventDTO(&event))
		} else if state == "completed" && event.Date.Before(time.Now()) {
			events = append(events, mappers_dto.EventDomainToEventDTO(&event))
		}
	}
	return events, nil
}
