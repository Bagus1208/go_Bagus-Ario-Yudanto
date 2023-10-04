package mock

import (
	"clean_architecture/features/users/entity"
	"errors"
)

type Repository struct {
	Status string
}

func (_m *Repository) Insert(data entity.User) (entity.User, error) {
	if _m.Status == "valid" {
		return data, nil
	}
	return entity.User{}, errors.New("error")
}

func (_m *Repository) GetAll() ([]entity.User, error) {
	if _m.Status == "valid" {
		listUser := []entity.User{
			{Email: "bagus@gmail.com", Password: "bagus123"},
			{Email: "rio@gmail.com", Password: "rio123"},
		}
		return listUser, nil
	}
	return nil, errors.New("error")
}

func (_m *Repository) Login(email string, password string) *entity.User {
	if _m.Status == "valid" {
		user := entity.User{
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	}
	return nil
}
