package input

type UserAddDTO struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=2"`
}
