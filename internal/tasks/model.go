package tasks

import "gorm.io/gorm"

type Task struct {
	*gorm.Model
	Name   string `json:"name" validate:"required"`
	Done   bool   `json:"done"`
	UserID uint   `json:"user_id"`
}

func NewTask(name string, id uint) *Task {
	return &Task{
		Name:   name,
		UserID: id,
	}
}
