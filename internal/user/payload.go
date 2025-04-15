package user

type DeleteRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ChangeRequest struct {
	NewName string `json:"new_name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}
