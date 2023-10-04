package repository

import (
	"clean_architecture/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data *model.User) (model.User, error)
	All() ([]model.User, error)
	Login(email string, password string) *model.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		db: DB,
	}
}

func (u *userRepository) Create(data *model.User) (model.User, error) {
	err := u.db.Create(data).Error
	if err != nil {
		return model.User{}, err
	}

	return *data, nil
}

func (u *userRepository) All() ([]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (u *userRepository) Login(email string, password string) *model.User {
	var user model.User
	if err := u.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil
	}
	if user.ID < 1 {
		return &user
	}

	return &user
}
