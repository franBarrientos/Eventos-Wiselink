package mappers_dto

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/domain/dtos/input"
	"github.com/franBarrientos/domain/dtos/output"
)

func EventDtoToEventDomain(event *input.EventAddDTO) *domain.Event {

	return &domain.Event{
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		LongDescription:  event.LongDescription,
		Date:             *event.Date,
		Organizer: domain.Organizer{
			FirstName: event.Organizer.FirstName,
			LastName:  event.Organizer.LastName,
		},
		Place: domain.Place{
			Country:       event.Place.Country,
			City:          event.Place.City,
			Address:       event.Place.Address,
			AddressNumber: event.Place.AddressNumber,
		},
		State: *event.State,
	}
}

func EventDomainToEventDTO(event *domain.Event) output.EventDTO {
	return output.EventDTO{
		Id:               event.Id,
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		LongDescription:  event.LongDescription,
		Date:             event.Date,
		Organizer: output.OrganizerDTO{
			Id:        event.Organizer.Id,
			FirstName: event.Organizer.FirstName,
			LastName:  event.Organizer.LastName,
		},
		Place: output.PlaceDTO{
			Id:            event.Place.Id,
			Country:       event.Place.Country,
			City:          event.Place.City,
			Address:       event.Place.Address,
			AddressNumber: event.Place.AddressNumber,
		},
		State: event.State,
	}
}
