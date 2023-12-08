package middleware

import (
	"github.com/franBarrientos/domain"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type JwtMiddleware func(c *fiber.Ctx) error

func (j JwtMiddleware) Serve(c *fiber.Ctx) error {
	return j(c)
}

func JwtAuthMiddleware(role string, jwtService domain.ITokenService) JwtMiddleware {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := jwtService.IsAuthorized(authToken)
			if authorized {
				roleExtracted, idExtracted, err := jwtService.ExtractRoleAndIDFromToken(authToken)

				if err != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"error": error.Error(err)})
				}

				if roleExtracted != role {
					return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{"error": "Forbidden you does not have access to this resource"})
				}

				c.Set("x-user-id", idExtracted)
				return c.Next()
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": "Please specify token in Authorization header"})
	}
}
