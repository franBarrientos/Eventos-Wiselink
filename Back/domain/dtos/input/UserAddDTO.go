package input

type UserAddDTO struct {
	FirstName string `validate:"required,max=40"`
	LastName  string `validate:"required,max=40"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=2"`
}
