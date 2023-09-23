package controller

import (
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/helper"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model model.UserModel
}

func (userController *UserController) InitUserController(userModel model.UserModel) {
	userController.model = userModel
}

func (usercontroller *UserController) GetUsersController(c echo.Context) error {
	var result = usercontroller.model.GetUsers()

	return c.JSON(http.StatusOK, helper.SetResponse("success get all users", result))
}

func (usercontroller *UserController) GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	var result = usercontroller.model.GetUser(id)

	return c.JSON(http.StatusOK, helper.SetResponse("success get user by id", result))
}

func (usercontroller *UserController) CreateUserController(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse(err.Error(), nil))
	}

	var result = usercontroller.model.CreateUser(user)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
	}

	return c.JSON(http.StatusOK, helper.SetResponse("success create user", result))
}

func (usercontroller *UserController) UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse(err.Error(), nil))
	}

	user.ID = uint(id)
	var result = usercontroller.model.UpdateUser(user)

	return c.JSON(http.StatusOK, helper.SetResponse("success update user by id", result))
}

func (usercontroller *UserController) DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	usercontroller.model.DeleteUser(id)

	return c.JSON(http.StatusOK, []any{})
}
