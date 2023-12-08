package output

type AuthResponse struct {
	User  UserDTO
	Token LoginResponse
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}
