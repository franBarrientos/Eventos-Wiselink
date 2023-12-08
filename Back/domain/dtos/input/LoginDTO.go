package input

type LoginDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=2"`
}
