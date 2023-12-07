package entities_db

import "time"

type Event struct {
	ID               int `gorm:"primaryKey;autoIncrement"`
	Title            string
	ShortDescription string
	LongDescription  string
	Date             time.Time
	OrganizerID      int
	Organizer        Organizer `gorm:"foreignKey:OrganizerID"`
	PlaceID          int
	Place            Place `gorm:"foreignKey:PlaceID"`
	State            bool
}
