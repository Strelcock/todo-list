package user

import (
	"todoProject/internal/models"
	"todoProject/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepositiry(db *db.Db) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	result := ur.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := ur.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) Delete(user *models.User) error {
	result := ur.Database.DB.Delete(user, "email = ?", user.Email)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) ChangeName(user *models.User, name string) error {
	result := ur.Database.DB.Model(user).Update("name", name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
