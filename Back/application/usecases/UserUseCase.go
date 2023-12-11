package usecases

import (
	"github.com/franBarrientos/application/mappers_dto"
	"github.com/franBarrientos/application/repositories"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/output"
)

type UserUseCase struct {
	userRepository repositories.IUserRepository
}

func NewUserUseCase(userRepository repositories.IUserRepository) domain.IUserUseCase {

	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u UserUseCase) GetEventsSubscribed(idUser int, state string, page int, limit int) ([]output.EventDTO, error) {

	eventsDomain, err := u.userRepository.GetEventsSubscribed(idUser, state, page, limit)
	if err != nil {
		return nil, err
	}
	events := []output.EventDTO{}
	for _, event := range eventsDomain {
		events = append(events, mappers_dto.EventDomainToEventDTO(&event))

	}

	return events, nil
}
