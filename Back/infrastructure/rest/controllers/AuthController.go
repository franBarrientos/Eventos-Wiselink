package controllers

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/infrastructure/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthUseCase domain.IAuthUseCase
	Validator   *validator.Validate
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var request input.LoginDTO

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := ac.Validator.Struct(request); err != nil {
		errorMessages := utils.ParseValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessages})
	}

	authResponse, err := ac.AuthUseCase.LoginUser(&request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(authResponse)
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var request input.UserAddDTO
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := ac.Validator.Struct(request); err != nil {
		errorMessages := utils.ParseValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessages})
	}

	authResponse, err := ac.AuthUseCase.RegisterUser(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(authResponse)
}
