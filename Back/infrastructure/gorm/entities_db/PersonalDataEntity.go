package entities_db

type PersonalData struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
}
