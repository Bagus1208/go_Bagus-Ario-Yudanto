package repository

import (
	"clean_architecture/features/users/entity"

	"gorm.io/gorm"
)

type userModel struct {
	db *gorm.DB
}

func New(DB *gorm.DB) entity.Repository {
	return &userModel{
		db: DB,
	}
}

func (usermodel *userModel) Insert(data entity.User) (entity.User, error) {
	err := usermodel.db.Create(data).Error
	if err != nil {
		return entity.User{}, err
	}
	return data, nil
}

func (usermodel *userModel) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := usermodel.db.Find(&users).Error
	if err != nil {
		return []entity.User{}, err
	}
	return users, nil
}

func (usermodel *userModel) Login(email string, password string) *entity.User {
	var user entity.User
	if err := usermodel.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil
	}
	if user.ID < 1 {
		return &user
	}
	return &user
}
