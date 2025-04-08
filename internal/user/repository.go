package user

import "todoProject/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepositiry(db *db.Db) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (ur *UserRepository) Create(user *User) (*User, error) {
	result := ur.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := ur.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
