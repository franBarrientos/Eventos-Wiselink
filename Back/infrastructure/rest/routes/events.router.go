package routes

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/infrastructure/rest/controllers"
	middleware "github.com/franBarrientos/infrastructure/rest/middlewere"
	"github.com/gofiber/fiber/v2"
)

func EventsRoutes(router fiber.Router, controller *controllers.EventController, jwtService *domain.ITokenService) {

	authGroup := router.Group("/admin", middleware.JwtAuthMiddleware("ADMIN", *jwtService))

	authGroup.Get("/events", controller.GetAllEvents)
	authGroup.Post("/events", controller.CreateEvent)
	authGroup.Put("/events/:id", controller.UpdateEvent)
	router.Get("/events", controller.GetEventsFiltered)
	router.Post("/events", controller.SubscribeUserToEvent)
	router.Get("/events/user/:id", controller.GetEventsSubscribedUser)

}
