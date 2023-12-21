package routes

import (
	"github.com/franBarrientos/application/usecases"
	"github.com/franBarrientos/infrastructure/config"
	"github.com/franBarrientos/infrastructure/gorm/repositories_db"
	"github.com/franBarrientos/infrastructure/rest/controllers"
	middleware "github.com/franBarrientos/infrastructure/rest/middlewere"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB, env *config.Env) {

	eventRepository := repositories_db.NewEventRepositoryDb(db)
	userRepository := repositories_db.NewUserRepositoryDb(db)

	eventUseCase := usecases.NewEventUseCase(eventRepository)
	userUseCase := usecases.NewUserUseCase(userRepository)

	jwtService := middleware.CreateJwtService(env)
	authUseCase := usecases.NewAuthUseCase(userRepository, jwtService)

	validator := validator.New(validator.WithRequiredStructEnabled())

	eventController := controllers.EventController{
		EventUseCase: eventUseCase,
		Validator:    validator,
		UserUseCase:  userUseCase,
	}

	authController := controllers.AuthController{
		AuthUseCase: authUseCase,
		Validator:   validator,
	}

	eventsRoutes := app.Group("/api/v1")
	EventsRoutes(eventsRoutes, &eventController, &jwtService)

	authRoutes := app.Group("/auth")
	AuthRoutes(authRoutes, &authController)

}
