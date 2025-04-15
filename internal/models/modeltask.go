package models

import (
	"gorm.io/gorm"
)

// var errCantCreateTask = errors.New("cannot create task")

type Task struct {
	*gorm.Model
	Name string `json:"name" validate:"required"`
	Done bool   `json:"done"`

	UserID uint `json:"user_id" gorm:"foreignKey:UserID"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func NewTask(name string, id uint) *Task {
	return &Task{
		Name:   name,
		UserID: id,
	}
}
