package entities_db

import (
	"github.com/franBarrientos/domain"
)

type User struct {
	ID               int `gorm:"primaryKey;autoIncrement"`
	Email            string
	Password         string
	Role             domain.Role `gorm:"type:enum('ADMIN', 'USER')"`
	PersonalDataID   int
	PersonalData     PersonalData `gorm:"foreignKey:PersonalDataID"`
	EventsSubscribed []Event      `gorm:"many2many:user_events;"`
}
