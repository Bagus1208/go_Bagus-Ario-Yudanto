package controller

import (
	"REST_API_testing/configs"
	"REST_API_testing/helper"
	"REST_API_testing/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetBlogs(c echo.Context) error
	Login(c echo.Context) error
}

type UserController struct {
	model  model.UserModelInterface
	config configs.Config
}

func NewUserController(userModelInterface model.UserModelInterface, conf configs.Config) UserControllerInterface {
	return &UserController{
		model:  userModelInterface,
		config: conf,
	}
}

// func (userController *UserController) InitUserController(userModel model.UserModel, config configs.Config) {
// 	userController.model = userModel
// 	userController.config = config
// }

func (userController *UserController) GetUsers(c echo.Context) error {
	var result = userController.model.GetAllUser()

	return c.JSON(http.StatusOK, helper.SetResponse("success get all users", result))
}

func (userController *UserController) GetUser(c echo.Context) error {
	input := c.Param("id")
	id, err := strconv.Atoi(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse("Invalid ID", nil))
	}

	var result = userController.model.GetUserById(id)

	return c.JSON(http.StatusOK, helper.SetResponse("success get user", result))
}

func (userController *UserController) Create(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	var result = userController.model.CreateUser(user)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
	}

	return c.JSON(http.StatusCreated, helper.SetResponse("success create user", result))
}

func (userController *UserController) Update(c echo.Context) error {
	var update model.User
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse("Invalid ID", nil))
	}

	update.ID = uint(id)
	var result = userController.model.UpdateUser(update)

	return c.JSON(http.StatusOK, helper.SetResponse("success update", result))
}

func (userController *UserController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse("Invalid ID", nil))
	}

	userController.model.DeleteUser(id)

	return c.JSON(http.StatusOK, nil)
}

func (userController *UserController) GetBlogs(c echo.Context) error {
	var result = userController.model.GetUserBlogs()

	return c.JSON(http.StatusOK, helper.SetResponse("success get user's blogs", result))
}

func (usercontroller *UserController) Login(c echo.Context) error {
	var input = model.LoginModel{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
	}

	var result = usercontroller.model.Login(input.Email, input.Password)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
	}

	if result.ID < 1 {
		return c.JSON(http.StatusNotFound, helper.SetResponse("data not found", nil))
	}

	var jwtToken = helper.GenerateJWT(usercontroller.config.Secret, usercontroller.config.Secret, int(result.ID))
	if jwtToken == nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data", nil))
	}

	jwtToken["info"] = result

	return c.JSON(http.StatusOK, helper.SetResponse("success", jwtToken))
}
