package service

import (
	"clean_architecture/features/users/dto"
	"clean_architecture/features/users/entity"
	"log"

	"github.com/mashingan/smapping"
)

type userService struct {
	userRepo entity.Repository
}

func New(UserRepo entity.Repository) entity.Service {
	return &userService{
		userRepo: UserRepo,
	}
}

func (service *userService) Create(userdto dto.UserInput) (entity.User, error) {
	var user entity.User
	err := smapping.FillStruct(&user, smapping.MapFields(&userdto))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := service.userRepo.Insert(&user)
	if err != nil {
		return entity.User{}, err
	}
	return res, nil
}

func (service *userService) GetAll() ([]entity.User, error) {
	return service.userRepo.GetAll()
}

func (service *userService) Login(email string, password string) *entity.User {
	var user = new(entity.User)
	user = service.userRepo.Login(email, password)

	return user
}
