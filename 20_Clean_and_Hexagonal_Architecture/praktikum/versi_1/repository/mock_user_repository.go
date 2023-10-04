package repository

import (
	"clean_architecture/model"
	"errors"
)

type MockUserRepository struct {
	Status string
}

func (mockUser *MockUserRepository) Create(data *model.User) (model.User, error) {
	if mockUser.Status == "valid" {
		return *data, nil
	}
	return model.User{}, errors.New("error")
}

func (mockUser *MockUserRepository) All() ([]model.User, error) {
	if mockUser.Status == "valid" {
		listUser := []model.User{
			{Email: "bagus@gmail.com", Password: "bagus123"},
			{Email: "rio@gmail.com", Password: "rio123"},
		}
		return listUser, nil
	}
	return nil, errors.New("error")
}

func (mockUser *MockUserRepository) Login(email string, password string) *model.User {
	if mockUser.Status == "valid" {
		user := model.User{
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	}
	return nil
}
