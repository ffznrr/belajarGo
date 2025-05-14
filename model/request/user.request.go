package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type UserLoginRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}