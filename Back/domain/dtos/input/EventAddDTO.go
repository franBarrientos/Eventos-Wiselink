package input

import (
	"time"
)

type EventAddDTO struct {
	Title            string `validate:"required,min=5,max=40"`
	ShortDescription string `validate:"required,min=5,max=80"`
	LongDescription  string `validate:"required,min=5,max=200"`
	Date             *time.Time
	Organizer        *OrganizerAddDTO `validate:"required"`
	Place            *PlaceAddDTO     `validate:"required"`
	State            *bool            `validate:"required"`
}

/*
func (e *EventAddDTO) UnmarshalJSON(data []byte) error {

	type TempEvent struct {
		Title            string
		ShortDescription string
		LongDescription  string
		Date             string
		Organizer        *OrganizerAddDTO
		Place            *PlaceAddDTO
		State            *bool
	}

	// I need to use a temporary struct to unmarshal the data because the date received is a string
	tempEvent := TempEvent{}
	if err := json.Unmarshal(data, &tempEvent); err != nil {
		return err
	}

	e.Title = tempEvent.Title
	e.ShortDescription = tempEvent.ShortDescription
	e.LongDescription = tempEvent.LongDescription
	e.Organizer = tempEvent.Organizer
	e.Place = tempEvent.Place
	e.State = tempEvent.State

	parsedTime, err := time.Parse("2006-01-02T15:04", tempEvent.Date)
	if err != nil {
		return err
	}

	e.Date = &parsedTime

	return nil
}
*/
