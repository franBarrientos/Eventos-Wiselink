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
	page := utils.PageOrDefault(c.Query("page"), 1)
	limit := utils.LimitOrDefault(c.Query("limit"), 12)

	events, err := ec.EventUseCase.GetAllEvents(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(events)

}

func (ec *EventController) CreateEvent(c *fiber.Ctx) error {
	event := input.EventAddDTO{}

	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if err := ec.Validator.Struct(event); err != nil {
		errorMessages := utils.ParseValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessages})
	}

	eventCreated, err := ec.EventUseCase.CreateEvent(&event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(eventCreated)

}

func (ec *EventController) GetEventsFiltered(c *fiber.Ctx) error {
	title := c.Query("title")
	state := c.Query("state")
	date := c.Query("date")

	page := utils.PageOrDefault(c.Query("page"), 1)
	limit := utils.LimitOrDefault(c.Query("limit"), 12)

	result, err := ec.EventUseCase.GetEventsFiltered(date, state, title, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
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

	result, errorFromUpdated := ec.EventUseCase.UpdateEvent(int(id), &eventChanges)
	if errorFromUpdated != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errorFromUpdated.Error(),
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

	//if user and id from request params don't match, throw 403
	if strconv.Itoa(subscription.User) != c.Locals("UserId") {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	if err := ec.Validator.Struct(subscription); err != nil {
		errorMessages := utils.ParseValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessages})
	}

	errorSubscribe := ec.EventUseCase.SubscribeToEvent(&subscription)
	if errorSubscribe != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errorSubscribe.Error(),
		})
	}

	events, err := ec.UserUseCase.GetEventsSubscribed(subscription.User, "", 1, 20)
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

	page := utils.PageOrDefault(c.Query("page"), 1)
	limit := utils.LimitOrDefault(c.Query("limit"), 12)

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	//if user and id from request params don't match, throw 403
	if c.Params("id") != c.Locals("UserId") {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	events, errorFromGet := ec.UserUseCase.GetEventsSubscribed(int(id), c.Query("state"), page, limit)
	if errorFromGet != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errorFromGet.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"eventsSubscribed": events,
	})
}

func (ec *EventController) GetSubscribersToEvent(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	page := utils.PageOrDefault(c.Query("page"), 1)
	limit := utils.LimitOrDefault(c.Query("limit"), 12)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	users, errorFromGet := ec.EventUseCase.GetSubscribersToEvent(int(id), page, limit)
	if errorFromGet != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errorFromGet.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)

}
