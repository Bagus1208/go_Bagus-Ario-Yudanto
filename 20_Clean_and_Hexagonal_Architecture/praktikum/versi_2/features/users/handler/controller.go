package handler

import (
	"clean_architecture/config"
	"clean_architecture/features/users/dto"
	"clean_architecture/features/users/entity"
	"clean_architecture/helper"
	"clean_architecture/middleware"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService entity.Service
	config      config.Config
}

func New(UserService entity.Service, Config config.Config) entity.Handler {
	return &userController{
		userService: UserService,
		config:      Config,
	}
}

func (controll *userController) Create() echo.HandlerFunc {
	userInput := dto.UserInput{}
	return func(c echo.Context) error {
		if err := c.Bind(&userInput); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
		}
		user, err := controll.userService.Create(userInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success create user", user))
	}
}

func (controll *userController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := controll.userService.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success get all users", users))
	}
}

func (controll *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = dto.LoginInput{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
		}

		user := controll.userService.Login(input.Email, input.Password)
		if user == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
		}

		if user.ID < 1 {
			return c.JSON(http.StatusNotFound, helper.SetResponse("data not found", nil))
		}

		var jwtToken = middleware.GenerateJWT(controll.config.Secret, controll.config.Secret, int(user.ID))
		if jwtToken == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data", nil))
		}

		jwtToken["info"] = user

		return c.JSON(http.StatusOK, helper.SetResponse("success", jwtToken))
	}
}
