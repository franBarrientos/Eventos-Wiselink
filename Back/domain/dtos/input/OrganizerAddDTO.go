package input

type OrganizerAddDTO struct {
	FirstName string `validate:"required,max=40"`
	LastName  string `validate:"required,max=40"`
}
