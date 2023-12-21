package mappers_dto

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
)

func UserDomainToUserDTO(user *domain.User) output.UserDTO {

	var eventsDto []output.EventDTO
	for _, event := range user.Events {
		eventsDto = append(eventsDto, EventDomainToEventDTO(&event))
	}

	return output.UserDTO{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Events:    eventsDto,
	}
}

func UserAddDTOToUserDomain(user *input.UserAddDTO) domain.User {
	return domain.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      string(domain.ADMIN),
	}
}
