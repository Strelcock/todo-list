package di

import (
	"todoProject/internal/models"
)

type IUserRepository interface {
	Create(*models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Delete(*models.User) error
}
