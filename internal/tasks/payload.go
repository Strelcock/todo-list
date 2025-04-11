package tasks

type CreateRequest struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}
