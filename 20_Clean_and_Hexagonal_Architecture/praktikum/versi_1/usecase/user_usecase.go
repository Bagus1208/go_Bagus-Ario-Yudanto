package usecase

import (
	"clean_architecture/dto"
	"clean_architecture/model"
	"clean_architecture/repository"
	"log"

	"github.com/mashingan/smapping"
)

type UserUseCase interface {
	Create(userdto dto.UserDTO) (model.User, error)
	All() ([]model.User, error)
	Login(email string, password string) *model.User
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepo,
	}
}

func (usecase *userUseCase) Create(userdto dto.UserDTO) (model.User, error) {
	var user model.User
	err := smapping.FillStruct(&user, smapping.MapFields(&userdto))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := usecase.userRepository.Create(&user)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (usecase *userUseCase) All() ([]model.User, error) {
	return usecase.userRepository.All()
}

func (usercase *userUseCase) Login(email string, password string) *model.User {
	var user = new(model.User)
	user = usercase.userRepository.Login(email, password)

	return user
}
