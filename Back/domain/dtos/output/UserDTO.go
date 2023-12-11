package output

type UserDTO struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Role      string
	Events    []EventDTO
}
