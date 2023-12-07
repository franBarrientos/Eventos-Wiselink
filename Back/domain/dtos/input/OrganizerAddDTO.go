package input

type OrganizerAddDTO struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
}
