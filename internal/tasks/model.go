package tasks

import (
	"todoProject/internal/user"

	"gorm.io/gorm"
)

type Task struct {
	*gorm.Model
	Name string `json:"name" validate:"required"`
	Done bool   `json:"done"`

	UserID uint      `json:"user_id" gorm:"foreignKey:UserID"`
	User   user.User `json:"-" gorm:"constraint:OnUodate:CASCADE,OnDelete:CASCADE"`
}

func NewTask(name string, id uint) *Task {
	return &Task{
		Name:   name,
		UserID: id,
	}
}
