package entity

import (
	"clean_architecture/features/users/dto"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	Insert(data User) (User, error)
	GetAll() ([]User, error)
	Login(email string, password string) *User
}

type Service interface {
	Create(userdto dto.UserInput) (User, error)
	GetAll() ([]User, error)
	Login(email string, password string) *User
}

type Handler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Login() echo.HandlerFunc
}
