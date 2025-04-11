package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

type DeleteRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	JWT string `json:"jwt"`
}

type LoginResponse struct {
	JWT string `json:"jwt"`
}
