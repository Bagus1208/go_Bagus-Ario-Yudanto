package controller

import (
	"fmt"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_eksplorasi/helper"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_eksplorasi/model"
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

func (userController *UserController) GetUsers(c echo.Context) error {
	var result = userController.model.GetAllUser()

	return c.JSON(http.StatusOK, helper.SetResponse("success get all users", result))
}

func (userController *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
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
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	update.ID = uint(id)
	var result = userController.model.UpdateUser(update)

	return c.JSON(http.StatusOK, helper.SetResponse("success update", result))
}

func (userController *UserController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	userController.model.DeleteUser(id)

	return c.JSON(http.StatusOK, []any{})
}

func (userController *UserController) GetBlogs(c echo.Context) error {
	var result = userController.model.GetUserBlogs()

	return c.JSON(http.StatusOK, helper.SetResponse("success get user's blogs", result))
}
