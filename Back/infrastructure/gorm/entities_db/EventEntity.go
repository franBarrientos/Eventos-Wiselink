package entities_db

import "time"

type Event struct {
	ID               int    `gorm:"primaryKey;autoIncrement"`
	Title            string `gorm:"size:40;not null"`
	ShortDescription string `gorm:"size:80;not null"`
	LongDescription  string `gorm:"size:200;not null"`
	Date             time.Time
	OrganizerID      int
	Organizer        Organizer `gorm:"foreignKey:OrganizerID"`
	PlaceID          int
	Place            Place `gorm:"foreignKey:PlaceID"`
	State            bool
}
