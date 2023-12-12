package repositories_db

import (
	"errors"
	"fmt"
	"github.com/franBarrientos/application/repositories"
	"github.com/franBarrientos/domain"
	"github.com/franBarrientos/infrastructure/gorm/entities_db"
	"github.com/franBarrientos/infrastructure/gorm/mappers_db"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type EventRepositoryDb struct {
	database *gorm.DB
}

func NewEventRepositoryDb(db *gorm.DB) repositories.IEventRepository {
	return &EventRepositoryDb{
		database: db,
	}
}

func (ev EventRepositoryDb) GetAllEvents(page int, limit int) ([]domain.Event, error) {
	var events []entities_db.Event
	result := ev.database.Offset((page - 1) * limit).
		Limit(limit).
		Preload("Organizer").
		Preload("Organizer.PersonalData").
		Preload("Place").
		Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}

	var dtosEvents []domain.Event
	for _, dbEvent := range events {
		dtosEvents = append(dtosEvents, mappers_db.EventEntityToEventDomain(&dbEvent))
	}
	return dtosEvents, nil
}

func (ev EventRepositoryDb) GetEventById(id int) (domain.Event, error) {
	var event entities_db.Event
	result := ev.database.Find(&event, id)
	if result.Error != nil {
		return domain.Event{}, result.Error
	}

	return mappers_db.EventEntityToEventDomain(&event), nil
}

func (ev EventRepositoryDb) GetEventsFiltered(date string, state string, title string,
	page int, limit int) ([]domain.Event, error) {
	var events []entities_db.Event
	query := ev.database.Model(&entities_db.Event{})
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if state == "active" {
		query = query.Where("date > NOW()")
	}
	if state == "completed" {
		query = query.Where("date < NOW()")
	}
	if date != "" {
		recivedDate, err := time.Parse(time.RFC3339, date)
		if err != nil {
			fmt.Println("Error al parsear la fecha recibida:", err)
			return nil, err
		}
		truncatedRecivedDate := recivedDate.Truncate(24 * time.Hour)

		formattedDate := truncatedRecivedDate.Format("2006-01-02 15:04:05.999")

		query = query.Where("DATE(date) = DATE(?)", formattedDate)

	}
	if err := query.Offset((page-1)*limit).
		Limit(limit).
		Where(" state = ?", true).
		Preload("Organizer").
		Preload("Organizer.PersonalData").
		Preload("Place").
		Find(&events).Error; err != nil {
		return nil, err
	}

	dtosEvents := []domain.Event{}
	for _, dbEvent := range events {
		dtosEvents = append(dtosEvents, mappers_db.EventEntityToEventDomain(&dbEvent))
	}
	return dtosEvents, nil
}

func (ev EventRepositoryDb) CreateEvent(e *domain.Event) (domain.Event, error) {
	evenEntity := mappers_db.EventDomainToEventEntity(e)
	result := ev.database.Create(&evenEntity)
	if result.Error != nil {
		return domain.Event{}, result.Error
	}

	return mappers_db.EventEntityToEventDomain(&evenEntity), nil

}

func (ev EventRepositoryDb) UpdateEvent(id int, e map[string]interface{}) (domain.Event, error) {
	var eventToUpdate entities_db.Event
	result := ev.database.Preload("Organizer").
		Preload("Organizer.PersonalData").
		Preload("Place").
		First(&eventToUpdate, id)
	if result.Error != nil {
		return domain.Event{}, result.Error
	}

	if e["Title"] != nil && e["Title"] != "" {
		eventToUpdate.Title = e["Title"].(string)
	}

	if e["LongDescription"] != nil && e["LongDescription"] != "" {
		eventToUpdate.LongDescription = e["LongDescription"].(string)
	}
	if e["ShortDescription"] != nil && e["ShortDescription"] != "" {
		eventToUpdate.ShortDescription = e["ShortDescription"].(string)
	}

	if e["Date"] != nil {
		parsedTime, err := time.Parse("2006-01-02T15:04", e["Date"].(string))
		if err != nil {
			fmt.Println("Error al parsear la fecha:", err)
			return domain.Event{}, err

		}

		eventToUpdate.Date = parsedTime
	}

	if e["Organizer"] != nil {
		eventToUpdate.Organizer.PersonalData.FirstName = e["Organizer"].(map[string]interface{})["FirstName"].(string)
		eventToUpdate.Organizer.PersonalData.LastName = e["Organizer"].(map[string]interface{})["LastName"].(string)
		ev.database.Save(&eventToUpdate.Organizer.PersonalData)
	}

	if e["Place"] != nil {
		eventToUpdate.Place.Country = e["Place"].(map[string]interface{})["Country"].(string)
		eventToUpdate.Place.City = e["Place"].(map[string]interface{})["City"].(string)
		eventToUpdate.Place.Address = e["Place"].(map[string]interface{})["Address"].(string)
		addressNumber := e["Place"].(map[string]interface{})["AddressNumber"].(string)
		intValue, err := strconv.Atoi(addressNumber)
		if err != nil {
			return domain.Event{}, err
		}
		eventToUpdate.Place.AddressNumber = intValue
		ev.database.Save(&eventToUpdate.Place)
	}

	if e["State"] != nil {
		eventToUpdate.State = e["State"].(bool)
	}

	resultUpdate := ev.database.Save(&eventToUpdate)
	if resultUpdate.Error != nil {
		return domain.Event{}, resultUpdate.Error
	}

	return mappers_db.EventEntityToEventDomain(&eventToUpdate), nil

}

func (ev EventRepositoryDb) AddSubscribe(subscribe int, event int) error {

	var eventToSubscribe entities_db.Event
	result := ev.database.Where("date > NOW() and state = ? and id = ?", true, event).
		First(&eventToSubscribe, event)
	if result.Error != nil {
		return result.Error
	}

	var userToSubscribe entities_db.User
	result = ev.database.Preload("EventsSubscribed").First(&userToSubscribe, subscribe)
	if result.Error != nil {
		return result.Error
	}

	// Verificar si el evento ya está suscrito por el usuario
	eventAlreadySubscribed := false
	for _, subscribedEvent := range userToSubscribe.EventsSubscribed {
		if subscribedEvent.ID == event {
			eventAlreadySubscribed = true
			break
		}
	}

	if eventAlreadySubscribed {
		return errors.New("event already subscribed")
	}

	userToSubscribe.EventsSubscribed = append(userToSubscribe.EventsSubscribed, eventToSubscribe)
	ev.database.Save(&userToSubscribe)
	return nil
}

func (ev EventRepositoryDb) GetSubscribersToEvent(eventId int, page int, limit int) ([]domain.User, error) {

	subscribersEntity := []entities_db.User{}
	result := ev.database.
		Offset((page-1)*limit).
		Limit(limit).
		Joins("JOIN user_events ON users.id = user_events.user_id").
		Where("user_events.event_id = ?", eventId).
		Preload("EventsSubscribed").
		Preload("EventsSubscribed.Place").
		Preload("EventsSubscribed.Organizer").
		Preload("PersonalData").
		Find(&subscribersEntity)
	if result.Error != nil {
		return nil, result.Error
	}

	subscribersDomain := []domain.User{}
	for _, subscriber := range subscribersEntity {
		subscribersDomain = append(subscribersDomain, mappers_db.UserEntityToUserDomain(&subscriber))
	}
	return subscribersDomain, nil
}
