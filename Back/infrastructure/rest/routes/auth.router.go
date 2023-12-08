package routes

import (
	"github.com/franBarrientos/infrastructure/rest/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, controller *controllers.AuthController) {

	router.Post("/login", controller.Login)
	router.Post("/register", controller.Register)
}
