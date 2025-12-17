package service

import (
	"errors"
	"go-ecommerce-api/internal/domain"
	"go-ecommerce-api/internal/dto"
	"go-ecommerce-api/internal/helper"
	"go-ecommerce-api/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {
	log.Println(input)

	hPassword, err := s.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		return "", err
	}
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	log.Println(user)

	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)

}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.FindUserByEmail(email)

	if err != nil {
		return "", errors.New("user does not exist with provided email")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)

}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}
