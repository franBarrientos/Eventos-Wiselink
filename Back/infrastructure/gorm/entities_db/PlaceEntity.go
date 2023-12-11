package entities_db

type Place struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	Address       string `gorm:"size:40;not null"`
	AddressNumber int    `gorm:"not null"`
	City          string `gorm:"size:40;not null"`
	Country       string `gorm:"size:40;not null"`
}
