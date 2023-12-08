package middleware

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	Name string
	ID   int
	Role string
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID int
	jwt.RegisteredClaims
}
