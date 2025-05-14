package response

type LoginResponse struct {
	Name  string
	Token string
}

type RegisterResponse struct {
	Name    string
	Email   string
	Address string
}