package output

type UserDTO struct {
	FirstName string
	LastName  string
	Email     string
	Role      string
	Events    []EventDTO
}
