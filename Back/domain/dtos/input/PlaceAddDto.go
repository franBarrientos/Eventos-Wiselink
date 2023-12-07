package input

type PlaceAddDTO struct {
	Address       string `validate:"required"`
	AddressNumber int    `validate:"required"`
	City          string `validate:"required"`
	Country       string `validate:"required"`
}
