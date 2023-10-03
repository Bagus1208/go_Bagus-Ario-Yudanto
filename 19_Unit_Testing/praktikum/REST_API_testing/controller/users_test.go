package controller

import (
	"REST_API_testing/configs"
	"REST_API_testing/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type UserMock struct {
	status string
}

func (userMock *UserMock) GetAllUser() []model.User {
	if userMock.status == "valid" {
		listUser := []model.User{
			{Name: "Bagus"}, {Name: "Ario"}, {Name: "Yudanto"},
		}
		return listUser
	}
	return nil
}

func (userMock *UserMock) GetUserById(id int) model.User {
	if userMock.status == "valid" {
		result := model.User{
			Model:    gorm.Model{ID: 1},
			Name:     "Bagus",
			Email:    "bagus@gmail.comn",
			Password: "bagus123",
		}
		return result
	}
	return model.User{}
}

func (userMock *UserMock) CreateUser(newUser model.User) *model.User {
	if userMock.status == "valid" {
		return &newUser
	}
	return nil
}

func (userMock *UserMock) UpdateUser(updateUser model.User) *model.User {
	if userMock.status == "valid" {
		return &updateUser
	}
	return nil
}

func (userMock *UserMock) DeleteUser(id int) {}

func (userMock *UserMock) GetUserBlogs() []model.User {
	if userMock.status == "valid" {
		listUser := []model.User{
			{Name: "Bagus"}, {Name: "Ario"}, {Name: "Yudanto"},
		}
		return listUser
	}
	return nil
}

func (userMock *UserMock) Login(email string, password string) *model.User {
	if userMock.status == "valid" {
		user := model.User{
			Model:    gorm.Model{ID: 1},
			Name:     "Bagus",
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	} else if userMock.status == "ID < 1" {
		user := model.User{
			Model:    gorm.Model{ID: 0},
			Name:     "Bagus",
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}
		return &user
	}
	return nil
}

type ResponseData struct {
	Data    map[string]any `json:"data"`
	Message string         `json:"message"`
}

var Token = "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzMDMzMzYsImlhdCI6MTY5NjI5OTczNiwiaWQiOjF9.TidWxlTxZyrzGGRYrgupD4DE6hVs_QTFmfsImT_tLDg"

func TestGetAllUsers(t *testing.T) {
	var mdl = UserMock{status: "valid"}
	var config = configs.InitConfig()
	var ctl = NewUserController(&mdl, *config)
	var e = echo.New()

	e.GET("/users", ctl.GetUsers, echojwt.JWT([]byte(config.Secret)))

	var req = httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderAuthorization, Token)
	var res = httptest.NewRecorder()

	e.ServeHTTP(res, req)

	type ResponseData struct {
		Data    []map[string]any `json:"data"`
		Message string           `json:"message"`
	}

	var tmp = ResponseData{}
	var resData = json.NewDecoder(res.Result().Body)
	err := resData.Decode(&tmp)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotNil(t, tmp.Data)
	assert.Equal(t, "success get all users", tmp.Message)
}

func TestGetUserById(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		var mdl = UserMock{status: "invalid"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.GET("/users/:id", ctl.GetUser, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodGet, "/users/satu", nil)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "Invalid ID", tmp.Message)
	})

	t.Run("Success", func(t *testing.T) {
		var mdl = UserMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.GET("/users/:id", ctl.GetUser, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodGet, "/users/1", nil)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success get user", tmp.Message)
	})
}

func TestCreateUser(t *testing.T) {
	var mdl = UserMock{status: "invalid"}
	var config = configs.InitConfig()
	var ctl = NewUserController(&mdl, *config)
	var e = echo.New()

	e.POST("/users", ctl.Create)

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"name":"Bagus", "email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPost, "/users", bodyReq)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "error when parshing data -code=415, message=Unsupported Media Type", tmp.Message)
	})

	t.Run("Server error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"name":"Bagus", "email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPost, "/users", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "cannot process data, something happend", tmp.Message)
	})

	t.Run("Success", func(t *testing.T) {
		var mdl = UserMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.POST("/users", ctl.Create)

		var bodyReq = bytes.NewReader([]byte(`{"name":"Bagus", "email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPost, "/users", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success create user", tmp.Message)
	})
}

func TestUpdateUser(t *testing.T) {
	var mdl = UserMock{status: "invalid"}
	var config = configs.InitConfig()
	var ctl = NewUserController(&mdl, *config)
	var e = echo.New()

	e.PUT("/users/:id", ctl.Update, echojwt.JWT([]byte(config.Secret)))

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"name":"Bagus", "email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPut, "/users/1", bodyReq)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "error when parshing data -code=415, message=Unsupported Media Type", tmp.Message)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"name":"Bagus", "email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPut, "/users/satu", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "Invalid ID", tmp.Message)
	})

	t.Run("Success", func(t *testing.T) {
		var mdl = UserMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.PUT("/users/:id", ctl.Update, echojwt.JWT([]byte(config.Secret)))

		var bodyReq = bytes.NewReader([]byte(`{"name":"Bagus", "email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPut, "/users/1", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success update", tmp.Message)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		var mdl = UserMock{status: "invalid"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.DELETE("/users/:id", ctl.Delete, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodDelete, "/users/satu", nil)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "Invalid ID", tmp.Message)
	})

	t.Run("Success", func(t *testing.T) {
		var mdl = UserMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.DELETE("/users/:id", ctl.Delete, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Nil(t, tmp.Data)
	})
}

func TestGetBlogs(t *testing.T) {
	var mdl = UserMock{status: "valid"}
	var config = configs.InitConfig()

	var ctl = NewUserController(&mdl, *config)

	var e = echo.New()
	e.GET("/users", ctl.GetUsers, echojwt.JWT([]byte(config.Secret)))

	var req = httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderAuthorization, Token)
	var res = httptest.NewRecorder()

	e.ServeHTTP(res, req)

	type ResponseData struct {
		Data    []map[string]any `json:"data"`
		Message string           `json:"message"`
	}

	var tmp = ResponseData{}
	var resData = json.NewDecoder(res.Result().Body)
	err := resData.Decode(&tmp)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotNil(t, tmp.Data)
	assert.Equal(t, "success get all users", tmp.Message)
}

func TestLogin(t *testing.T) {
	var mdl = UserMock{status: "invalid"}
	var config = configs.InitConfig()
	var ctl = NewUserController(&mdl, *config)
	var e = echo.New()

	e.POST("/users/login", ctl.Login)

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPost, "/users/login", bodyReq)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "invalid user input", tmp.Message)
	})

	t.Run("Server error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPost, "/users/login", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "cannot process data, something happend", tmp.Message)
	})

	t.Run("ID < 1", func(t *testing.T) {
		var mdl = UserMock{status: "ID < 1"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.POST("/users/login", ctl.Login)

		var bodyReq = bytes.NewReader([]byte(`{"email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPost, "/users/login", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Equal(t, "data not found", tmp.Message)
	})

	t.Run("Success", func(t *testing.T) {
		var mdl = UserMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewUserController(&mdl, *config)
		var e = echo.New()

		e.POST("/users/login", ctl.Login)

		var bodyReq = bytes.NewReader([]byte(`{"email":"bagus@gmail.com", "password":"bagus123"}`))
		var req = httptest.NewRequest(http.MethodPost, "/users/login", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		fmt.Println("Access Token :", tmp.Data["access_token"])

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success", tmp.Message)
	})
}
