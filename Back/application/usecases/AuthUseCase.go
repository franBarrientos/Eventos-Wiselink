package usecases

import (
	"errors"
	"github.com/franBarrientos/application/mappers_dto"
	"github.com/franBarrientos/application/repositories"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepository repositories.IUserRepository
	tokenService   domain.ITokenService
}

func NewAuthUseCase(userRepository repositories.IUserRepository, tokenService domain.ITokenService) domain.IAuthUseCase {

	return &AuthUseCase{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (u AuthUseCase) RegisterUser(user *input.UserAddDTO) (output.AuthResponse, error) {
	passwordHashed, err := hashPassword(user.Password)
	if err != nil {
		return output.AuthResponse{}, err
	}

	user.Password = passwordHashed

	userDomain := mappers_dto.UserAddDTOToUseDomain(user)
	result, errFromCreate := u.userRepository.CreateUser(&userDomain)
	if errFromCreate != nil {
		return output.AuthResponse{}, errFromCreate
	}

	accessToken, _ := u.tokenService.CreateAccessToken(&result)
	refreshToken, _ := u.tokenService.CreateRefreshToken(&result)
	return output.AuthResponse{
		User: mappers_dto.UserDomainToUserDTO(&result),
		Token: output.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil

}

func (u AuthUseCase) LoginUser(credentials *input.LoginDTO) (output.AuthResponse, error) {

	result, err := u.userRepository.GetUserByEmail(credentials.Email)
	if err != nil {
		return output.AuthResponse{}, err
	}

	if !checkPassword(result.Password, credentials.Password) {
		return output.AuthResponse{}, errors.New("credentials are not valid")
	}

	accessToken, _ := u.tokenService.CreateAccessToken(&result)
	refreshToken, _ := u.tokenService.CreateRefreshToken(&result)

	return output.AuthResponse{
		User: mappers_dto.UserDomainToUserDTO(&result),
		Token: output.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil

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
