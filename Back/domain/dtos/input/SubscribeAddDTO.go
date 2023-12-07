package input

type SubscribeAddDTO struct {
	User  int `validate:"required"`
	Event int `validate:"required"`
}
