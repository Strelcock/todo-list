package di

import (
	"todoProject/internal/user"
)

type IUserRepository interface {
	Create(*user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
	Delete(*user.User) error
}
