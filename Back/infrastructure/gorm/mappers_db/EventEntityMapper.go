package mappers_db

import (
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/infrastructure/gorm/entities_db"
)

func EventEntityToEventDomain(event *entities_db.Event) domain.Event {
	return domain.Event{
		Id:               event.ID,
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		LongDescription:  event.LongDescription,
		Date:             event.Date,
		Organizer: domain.Organizer{
			Id:        event.OrganizerID,
			FirstName: event.Organizer.PersonalData.FirstName,
			LastName:  event.Organizer.PersonalData.LastName,
		},
		Place: domain.Place{
			Id:            event.PlaceID,
			Country:       event.Place.Country,
			City:          event.Place.City,
			AddressNumber: event.Place.AddressNumber,
			Address:       event.Place.Address,
		},
		State: event.State,
	}
}

func EventDomainToEventEntity(event *domain.Event) entities_db.Event {
	return entities_db.Event{
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		LongDescription:  event.LongDescription,
		Date:             event.Date,
		Organizer: entities_db.Organizer{
			PersonalData: entities_db.PersonalData{
				FirstName: event.Organizer.FirstName,
				LastName:  event.Organizer.LastName,
			},
		},
		Place: entities_db.Place{
			AddressNumber: event.Place.AddressNumber,
			Address:       event.Place.Address,
			Country:       event.Place.Country,
			City:          event.Place.City,
		},
		State: event.State,
	}
}
