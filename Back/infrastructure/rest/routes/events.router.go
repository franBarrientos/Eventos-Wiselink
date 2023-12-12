package routes

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/infrastructure/rest/controllers"
	middleware "github.com/franBarrientos/infrastructure/rest/middlewere"
	"github.com/gofiber/fiber/v2"
)

func EventsRoutes(router fiber.Router, controller *controllers.EventController, jwtService *domain.ITokenService) {

	//endpoints for admin
	authGroup := router.Group("/admin", middleware.JwtAuthMiddleware("ADMIN", *jwtService))
	authGroup.Get("/events", controller.GetAllEvents)
	authGroup.Post("/events", controller.CreateEvent)
	authGroup.Put("/events/:id", controller.UpdateEvent)
	authGroup.Get("/subscribers/event/:id", controller.GetSubscribersToEvent)

	//endpoints for user
	router.Post("/events", middleware.JwtAuthMiddleware("USER", *jwtService), controller.SubscribeUserToEvent)
	router.Get("/events/user/:id", middleware.JwtAuthMiddleware("USER", *jwtService), controller.GetEventsSubscribedUser)

	// public
	router.Get("/events", controller.GetEventsFiltered)

}
