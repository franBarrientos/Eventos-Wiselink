package usecases

import (
	"errors"
	"github.com/franBarrientos/application/mappers_dto"
	"github.com/franBarrientos/application/repositories"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
	"golang.org/x/crypto/bcrypt"
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

func (u UserUseCase) RegisterUser(user *input.UserAddDTO) (output.AuthResponse, error) {
	passwordHashed, err := hashPassword(user.Password)
	if err != nil {
		return output.AuthResponse{}, err
	}

	user.Password = passwordHashed

	userDomain := mappers_dto.UserAddDTOToUseDomain(user)
	result, err := u.userRepository.CreateUser(&userDomain)
	if err != nil {
		return output.AuthResponse{}, err
	}

	return output.AuthResponse{
		User:  mappers_dto.UserDomainToUserDTO(&result),
		Token: "",
	}, nil

}

func (u UserUseCase) LoginUser(email, password string) (output.AuthResponse, error) {

	result, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return output.AuthResponse{}, err
	}

	if !checkPassword(result.Password, password) {
		return output.AuthResponse{}, errors.New("credentials are not valid")
	}

	return output.AuthResponse{
		User:  mappers_dto.UserDomainToUserDTO(&result),
		Token: "",
	}, nil

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

func hashPassword(password string) (string, error) {
	// Generar un hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPassword(hashedPassword, inputPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword)) == nil
}
