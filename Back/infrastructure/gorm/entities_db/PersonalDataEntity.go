package entities_db

type PersonalData struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"size:40;not null"`
	LastName  string `gorm:"size:40;not null"`
}
