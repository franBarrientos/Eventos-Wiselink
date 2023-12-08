package controllers

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/infrastructure/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type EventController struct {
	EventUseCase domain.IEventUseCase
	UserUseCase  domain.IUserUseCase
	Validator    *validator.Validate
}

func (ec *EventController) GetAllEvents(c *fiber.Ctx) error {

	events, error := ec.EventUseCase.GetAllEvents()
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(events)
	}
}

func (ec *EventController) CreateEvent(c *fiber.Ctx) error {
	event := input.EventAddDTO{}
	c.BodyParser(&event)
	if err := c.BodyParser(&event); err != nil {
		return err
	}
	if err := ec.Validator.Struct(event); err != nil {
		// Presentar mensajes de error de manera más amigable
		errorMessages := utils.ParseValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessages})
	}

	eventCreated, error := ec.EventUseCase.CreateEvent(&event)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(eventCreated)

}

func (ec *EventController) GetEventsFiltered(c *fiber.Ctx) error {
	title := c.Query("title")
	state := c.Query("state")
	date := c.Query("date")

	result, error := ec.EventUseCase.GetEventsFiltered(date, state, title)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)

}

func (ec *EventController) UpdateEvent(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	var eventChanges map[string]interface{}
	if errParser := c.BodyParser(&eventChanges); errParser != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errParser.Error(),
		})
	}

	result, error := ec.EventUseCase.UpdateEvent(int(id), &eventChanges)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)

}

func (ec *EventController) SubscribeUserToEvent(c *fiber.Ctx) error {
	var subscription input.SubscribeAddDTO

	if errParser := c.BodyParser(&subscription); errParser != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errParser.Error(),
		})
	}

	if err := ec.Validator.Struct(subscription); err != nil {
		errorMessages := utils.ParseValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessages})
	}

	errorSubscribe := ec.EventUseCase.SubscribeToEvent(&subscription)
	if errorSubscribe != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": error.Error(errorSubscribe),
		})
	}

	events, err := ec.UserUseCase.GetEventsSubscribed(subscription.User, "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"eventsSubscribed": events,
	})

}

func (ec *EventController) GetEventsSubscribedUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	events, err := ec.UserUseCase.GetEventsSubscribed(int(id), c.Query("state"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"eventsSubscribed": events,
	})
}
