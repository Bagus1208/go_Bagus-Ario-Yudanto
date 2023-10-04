package mock

import (
	"clean_architecture/features/users/dto"
	"clean_architecture/features/users/entity"
	"errors"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type Service struct {
	Status string
}

func (_m *Service) Create(userdto dto.UserInput) (entity.User, error) {
	if _m.Status == "valid" {
		var user = entity.User{}
		smapping.FillStruct(&user, smapping.MapFields(&userdto))
		return user, nil
	}
	return entity.User{}, errors.New("error")
}

func (_m *Service) All() ([]entity.User, error) {
	if _m.Status == "valid" {
		listUser := []entity.User{
			{Email: "bagus@gmail.com", Password: "bagus123"},
			{Email: "rio@gmail.com", Password: "rio123"},
		}
		return listUser, nil
	}
	return nil, errors.New("error")
}

func (_m *Service) Login(email string, password string) *entity.User {
	if _m.Status == "valid" {
		user := entity.User{
			Model:    gorm.Model{ID: 1},
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	} else if _m.Status == "ID < 1" {
		user := entity.User{
			Model:    gorm.Model{ID: 0},
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	}
	return nil
}
