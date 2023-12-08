package domain

import (
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
)

type IAuthUseCase interface {
	LoginUser(credential *input.LoginDTO) (output.AuthResponse, error)
	RegisterUser(user *input.UserAddDTO) (output.AuthResponse, error)
}

type ITokenService interface {
	CreateAccessToken(user *User) (string, error)
	CreateRefreshToken(user *User) (string, error)
	IsAuthorized(requestToken string) (bool, error)
	/*	ExtractIDFromToken(requestToken string) (string, error)
	 */ExtractRoleAndIDFromToken(requestToken string) (string, string, error)
}
