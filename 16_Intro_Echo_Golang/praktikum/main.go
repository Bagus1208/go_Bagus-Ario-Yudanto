package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

func SetResponse(message string, data any) map[string]any {
	if data == nil {
		return map[string]any{
			"message": message,
		}
	}

	return map[string]any{
		"message": message,
		"data":    data,
	}
}

// get all users
func GetUsersController(c echo.Context) error {
	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, SetResponse("data not found", nil))
	}
	return c.JSON(http.StatusOK, SetResponse("messages: success get all users", users))
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
	}

	if id < 0 || id > len(users) {
		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
	}

	user := users[id-1]
	return c.JSON(http.StatusOK, SetResponse("success get user", user))
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
	}

	if id < 0 || id > len(users) {
		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
	}

	index := id - 1
	users = append(users[:index], users[index+1:]...)
	return c.JSON(http.StatusOK, []int{})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
	}

	if id < 0 || id > len(users) {
		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
	}

	index := id - 1
	users[index].Name = user.Name
	users[index].Email = user.Email
	users[index].Password = user.Password

	return c.JSON(http.StatusOK, SetResponse("success update", users[index]))
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusCreated, SetResponse("success create user", user))
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
