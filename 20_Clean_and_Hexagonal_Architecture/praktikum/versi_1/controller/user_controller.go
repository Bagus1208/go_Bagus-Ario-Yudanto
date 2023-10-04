package controller

import (
	"clean_architecture/config"
	"clean_architecture/dto"
	"clean_architecture/helper"
	"clean_architecture/middleware"
	"clean_architecture/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetAllUsers() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
}

type userController struct {
	userusecase usecase.UserUseCase
	config      config.Config
}

func NewUserController(usercase usecase.UserUseCase, conf config.Config) UserController {
	return &userController{
		userusecase: usercase,
		config:      conf,
	}
}

func (uc *userController) CreateUser() echo.HandlerFunc {
	userDTO := dto.UserDTO{}
	return func(c echo.Context) error {
		if err := c.Bind(&userDTO); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
		}
		user, err := uc.userusecase.Create(userDTO)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success create user", user))
	}
}

func (uc *userController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.userusecase.All()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success get all users", users))
	}
}

func (uc *userController) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = dto.LoginModel{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
		}

		user := uc.userusecase.Login(input.Email, input.Password)
		if user == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
		}

		if user.ID < 1 {
			return c.JSON(http.StatusNotFound, helper.SetResponse("data not found", nil))
		}

		var jwtToken = middleware.GenerateJWT(uc.config.Secret, uc.config.Secret, int(user.ID))
		if jwtToken == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data", nil))
		}

		jwtToken["info"] = user

		return c.JSON(http.StatusOK, helper.SetResponse("success", jwtToken))
	}
}
