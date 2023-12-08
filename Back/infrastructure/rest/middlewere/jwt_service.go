package middleware

import (
	"fmt"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/infrastructure/config"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type JwtService struct {
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
}

func CreateJwtService(env *config.Env) domain.ITokenService {
	return &JwtService{
		AccessTokenExpiryHour:  env.AccessTokenExpiryHour,
		RefreshTokenExpiryHour: env.RefreshTokenExpiryHour,
		AccessTokenSecret:      env.AccessTokenSecret,
		RefreshTokenSecret:     env.RefreshTokenSecret,
	}
}

func (s *JwtService) CreateAccessToken(user *domain.User) (accessToken string, err error) {
	exp := jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.AccessTokenExpiryHour)))
	claims := JwtCustomClaims{
		Name: user.FirstName,
		ID:   user.Id,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.AccessTokenSecret))
	if err != nil {
		return "", err
	}
	return t, err
}

func (s *JwtService) CreateRefreshToken(user *domain.User) (refreshToken string, err error) {
	claimsRefresh := JwtCustomRefreshClaims{
		ID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.RefreshTokenExpiryHour))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(s.RefreshTokenSecret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func (s *JwtService) IsAuthorized(requestToken string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.AccessTokenSecret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
	func (s *JwtService) ExtractIDFromToken(requestToken string) (string, error) {
		token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.AccessTokenSecret), nil
		})

		if err != nil {
			return "", err
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok && !token.Valid {
			return "", fmt.Errorf("Invalid Token")
		}

		return claims["ID"].(string), nil
	}
*/
func (s *JwtService) ExtractRoleAndIDFromToken(requestToken string) (string, string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.AccessTokenSecret), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", "", fmt.Errorf("Invalid Token")
	}
	idFloat, ok := claims["ID"].(float64)
	if !ok {
		return "", "", fmt.Errorf("ID field is not a valid float64")
	}

	id := strconv.FormatFloat(idFloat, 'f', -1, 64)

	return claims["Role"].(string), id, nil
}
