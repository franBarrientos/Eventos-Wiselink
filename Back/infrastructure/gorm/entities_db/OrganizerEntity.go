package entities_db

type Organizer struct {
	ID             int `gorm:"primaryKey;autoIncrement"`
	PersonalDataID int
	PersonalData   PersonalData `gorm:"foreignKey:PersonalDataID"`
}
