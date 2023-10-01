package test

import (
	"REST_API_testing/configs"
	"REST_API_testing/controller"
	"REST_API_testing/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type userResponse struct {
	Message string       `json:"message"`
	Data    []model.User `json:"data"`
}

func TestGetAllUsers(t *testing.T) {
	var config = configs.InitConfig()
	var DB = model.InitModel(*config)

	userModel := model.UserModel{}
	userModel.Init(DB)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userControll.GetUsers(c)

	var users userResponse
	body := rec.Body.String()
	err := json.Unmarshal([]byte(body), &users)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success get all users", users.Message)
}

func TestGetUser(t *testing.T) {
	var config = configs.InitConfig()
	var DB = model.InitModel(*config)

	userModel := model.UserModel{}
	userModel.Init(DB)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := userControll.GetUser(c)
	if err != nil {
		t.Error(err.Error())
	}

	var user struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}
	body := rec.Body.String()
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, uint(1), user.Data.ID)
	assert.Equal(t, "Bagus Yudanto", user.Data.Name)
	assert.Equal(t, "bagus@gmail.com", user.Data.Email)
}

func TestCreateUser(t *testing.T) {
	var userJSON = `{"name":"Reina","Email":"reina@gmail.com","password":"reina123"}`

	requestBody := strings.NewReader(userJSON)

	var config = configs.InitConfig()
	var DB = model.InitModel(*config)

	userModel := model.UserModel{}
	userModel.Init(DB)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", requestBody)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := userControll.Create(c)
	if err != nil {
		t.Error(err.Error())
	}

	var user struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}
	body := rec.Body.String()
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "Reina", user.Data.Name)
	assert.Equal(t, "reina@gmail.com", user.Data.Email)
	assert.Equal(t, "reina123", user.Data.Password)
}

func TestUpdateUser(t *testing.T) {
	var userJSON = `{"name":"Bagus Yudanto","Email":"bagus@gmail.com","password":"bagus123"}`

	requestBody := strings.NewReader(userJSON)

	var config = configs.InitConfig()
	var DB = model.InitModel(*config)

	userModel := model.UserModel{}
	userModel.Init(DB)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/:id", requestBody)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := userControll.Update(c)
	if err != nil {
		t.Error(err.Error())
	}

	var user struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}
	body := rec.Body.String()
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Bagus Yudanto", user.Data.Name)
	assert.Equal(t, "bagus@gmail.com", user.Data.Email)
	assert.Equal(t, "bagus123", user.Data.Password)
}

func TestDeleteUser(t *testing.T) {
	var config = configs.InitConfig()
	var DB = model.InitModel(*config)

	userModel := model.UserModel{}
	userModel.Init(DB)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := userControll.Delete(c)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, http.StatusOK, rec.Code)
}
