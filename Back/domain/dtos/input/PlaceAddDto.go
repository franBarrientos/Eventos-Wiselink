package input

type PlaceAddDTO struct {
	Address       string `validate:"required,max=40"`
	AddressNumber int    `validate:"required"`
	City          string `validate:"required,max=40"`
	Country       string `validate:"required,max=40"`
}
