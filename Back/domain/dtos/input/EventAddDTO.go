package input

import (
	"time"
)

type EventAddDTO struct {
	Title            string `validate:"required,min=2,max=40"`
	ShortDescription string `validate:"required,min=2,max=80"`
	LongDescription  string `validate:"required,min=2,max=200"`
	Date             *time.Time
	Organizer        *OrganizerAddDTO `validate:"required"`
	Place            *PlaceAddDTO     `validate:"required"`
	State            *bool            `validate:"required"`
}
