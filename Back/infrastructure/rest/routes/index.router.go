package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {

	eventsRoutes := app.Group("/api/v1")
	validator := validator.New(validator.WithRequiredStructEnabled())
	EventsRoutes(eventsRoutes, db, validator)
}
