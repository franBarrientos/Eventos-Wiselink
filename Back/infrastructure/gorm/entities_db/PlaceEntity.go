package entities_db

type Place struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	Address       string
	AddressNumber int
	City          string
	Country       string
}
