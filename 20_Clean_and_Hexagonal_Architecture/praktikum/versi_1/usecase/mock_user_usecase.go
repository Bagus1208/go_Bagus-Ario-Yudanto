package usecase

import (
	"clean_architecture/dto"
	"clean_architecture/model"
	"errors"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type MockUserUseCase struct {
	Status string
}

func (mockUser *MockUserUseCase) Create(userdto dto.UserDTO) (model.User, error) {
	if mockUser.Status == "valid" {
		var user = model.User{}
		smapping.FillStruct(&user, smapping.MapFields(&userdto))
		return user, nil
	}
	return model.User{}, errors.New("error")
}

func (mockUser *MockUserUseCase) All() ([]model.User, error) {
	if mockUser.Status == "valid" {
		listUser := []model.User{
			{Email: "bagus@gmail.com", Password: "bagus123"},
			{Email: "rio@gmail.com", Password: "rio123"},
		}
		return listUser, nil
	}
	return nil, errors.New("error")
}

func (mockUser *MockUserUseCase) Login(email string, password string) *model.User {
	if mockUser.Status == "valid" {
		user := model.User{
			Model:    &gorm.Model{ID: 1},
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	} else if mockUser.Status == "ID < 1" {
		user := model.User{
			Model:    &gorm.Model{ID: 0},
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	}
	return nil
}
