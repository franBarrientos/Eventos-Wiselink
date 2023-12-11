package usecases

import (
	"github.com/franBarrientos/application/mappers_dto"
	"github.com/franBarrientos/application/repositories"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
)

type EventUseCase struct {
	eventRepository repositories.IEventRepository
}

func (e EventUseCase) GetAllEvents(page int, limit int) ([]output.EventDTO, error) {

	events, err := e.eventRepository.GetAllEvents(page, limit)
	if err != nil {
		return nil, err
	}
	dtosEvents := []output.EventDTO{}
	for _, event := range events {
		dtosEvents = append(dtosEvents, mappers_dto.EventDomainToEventDTO(&event))
	}
	return dtosEvents, nil

}

func NewEventUseCase(eventRepository repositories.IEventRepository) domain.IEventUseCase {
	return &EventUseCase{
		eventRepository: eventRepository,
	}

}

func (e EventUseCase) GetEventsFiltered(date string, state string, title string, page int, limit int) ([]output.EventDTO, error) {

	events, err := e.eventRepository.GetEventsFiltered(date, state, title, page, limit)
	if err != nil {
		return nil, err
	}
	dtosEvents := []output.EventDTO{}
	for _, event := range events {
		dtosEvents = append(dtosEvents, mappers_dto.EventDomainToEventDTO(&event))
	}
	return dtosEvents, nil

}

func (e EventUseCase) CreateEvent(event *input.EventAddDTO) (output.EventDTO, error) {
	var eventEntity = mappers_dto.EventDtoToEventDomain(event)
	result, err := e.eventRepository.CreateEvent(eventEntity)

	if err != nil {
		return output.EventDTO{}, err
	} else {
		return mappers_dto.EventDomainToEventDTO(&result), nil
	}

}

func (e EventUseCase) UpdateEvent(id int, event *map[string]interface{}) (output.EventDTO, error) {

	userUpdated, err := e.eventRepository.UpdateEvent(id, *event)
	if err != nil {
		return output.EventDTO{}, err
	}
	return mappers_dto.EventDomainToEventDTO(&userUpdated), nil

}

func (e EventUseCase) SubscribeToEvent(subscribe *input.SubscribeAddDTO) error {

	err := e.eventRepository.AddSubscribe(subscribe.User, subscribe.Event)
	if err != nil {
		return err
	}
	return nil
}
