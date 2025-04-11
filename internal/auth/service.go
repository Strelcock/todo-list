package auth

import (
	"errors"
	"todoProject/internal/di"
	"todoProject/internal/user"

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

func (as *AuthService) Register(name, email, password string) (string, error) {
	foundUser, _ := as.UserRepo.FindByEmail(email)
	if foundUser != nil {
		return "", ErrUserExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	newUser := &user.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPass),
	}

	_, err = as.UserRepo.Create(newUser)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (as *AuthService) Login(email, password string) (string, error) {
	foundUser, _ := as.UserRepo.FindByEmail(email)
	if foundUser == nil {
		return "", ErrWrongCredentials
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return "", ErrWrongCredentials
	}

	return foundUser.Name, nil
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
