package output

import "time"

type EventDTO struct {
	Id               int
	Title            string
	ShortDescription string
	LongDescription  string
	Date             time.Time
	Organizer        OrganizerDTO
	Place            PlaceDTO
	State            bool
}
