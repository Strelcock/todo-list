package auth

import (
	"errors"
	"todoProject/internal/di"
	"todoProject/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo di.IUserRepository
}

var (
	ErrUserExists       = errors.New("user already exists")
	ErrWrongCredentials = errors.New("wrong email or password")
)

func NewAuthService(repo di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepo: repo,
	}
}

func (as *AuthService) Register(name, email, password string) (uint, error) {
	foundUser, _ := as.UserRepo.FindByEmail(email)
	if foundUser != nil {
		return 0, ErrUserExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	newUser := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPass),
	}

	_, err = as.UserRepo.Create(newUser)
	if err != nil {
		return 0, err
	}

	return newUser.ID, nil
}

func (as *AuthService) Login(email, password string) (uint, error) {
	foundUser, _ := as.UserRepo.FindByEmail(email)
	if foundUser == nil {
		return 0, ErrWrongCredentials
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return 0, ErrWrongCredentials
	}

	return foundUser.ID, nil
}

// func (as *AuthService) Delete(email, password string) error {
// 	foundUser, _ := as.UserRepo.FindByEmail(email)

// 	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
// 	if err != nil {
// 		return ErrWrongCredentials
// 	}

// 	err = as.UserRepo.Delete(foundUser)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
