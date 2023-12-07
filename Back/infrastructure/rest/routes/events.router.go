package routes

import (
	"github.com/franBarrientos/application/usecases"
	"github.com/franBarrientos/infrastructure/gorm/repositories_db"
	"github.com/franBarrientos/infrastructure/rest/controllers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func EventsRoutes(router fiber.Router, gorm *gorm.DB, validator *validator.Validate) {
	reporitory := repositories_db.NewEventRepositoryDb(gorm)
	userRepository := repositories_db.NewUserRepositoryDb(gorm)

	useCase := usecases.NewEventUseCase(reporitory)
	userUseCase := usecases.NewUserUseCase(userRepository)
	controller := controllers.EventController{
		EventUseCase: useCase,
		Validator:    validator,
		UserUseCase:  userUseCase,
	}

	router.Get("/admin/events", controller.GetAllEvents)
	router.Post("/admin/events", controller.CreateEvent)
	router.Put("/admin/events/:id", controller.UpdateEvent)
	router.Get("/events", controller.GetEventsFiltered)
	router.Post("/events", controller.SubscribeUserToEvent)
	router.Get("/events/user/:id", controller.GetEventsSubscribedUser)

}
