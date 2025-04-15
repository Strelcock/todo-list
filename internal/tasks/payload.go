package tasks

type CreateRequest struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

type MarkRequest struct {
	Status bool `json:"status"`
}

type ChangeNameRequest struct {
	NewName string `json:"new_name"`
}

type DeleteRequest struct {
	Name string `json:"name"`
}