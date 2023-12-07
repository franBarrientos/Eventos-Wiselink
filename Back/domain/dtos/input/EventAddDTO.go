package input

import "time"

type EventAddDTO struct {
	Title            string           `validate:"required"`
	ShortDescription string           `validate:"required,max=100"`
	LongDescription  string           `validate:"required,max=500"`
	Date             *time.Time       `validate:"required"`
	Organizer        *OrganizerAddDTO `validate:"required"`
	Place            *PlaceAddDTO     `validate:"required"`
	State            *bool            `validate:"required"`
}
