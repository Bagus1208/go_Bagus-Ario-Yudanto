package controller

import (
	"bytes"
	"clean_architecture/config"
	"clean_architecture/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type ResponseData struct {
	Data    map[string]any `json:"data"`
	Message string         `json:"message"`
}

var Token = "bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzNDMyMzUsImlhdCI6MTY5NjMzOTYzNSwiaWQiOjF9._8N1j56iF6BXJ0S_pbV9b4jz5ehwGXPLjSBg3pbjxoU"

func TestCreateUser(t *testing.T) {
	var mdl = usecase.MockUserUseCase{Status: "Invalid"}
	var conf = config.InitConfig()
	var ctl = NewUserController(&mdl, *conf)
	var e = echo.New()

	e.POST("/users", ctl.CreateUser())

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"email":"bagus@gmail.com", "password":"bagus123"}`))
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
		var bodyReq = bytes.NewReader([]byte(`{"email":"bagus@gmail.com", "password":"bagus123"}`))
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
		var mdl = usecase.MockUserUseCase{Status: "valid"}
		var conf = config.InitConfig()
		var ctl = NewUserController(&mdl, *conf)
		var e = echo.New()

		e.POST("/users", ctl.CreateUser())

		var bodyReq = bytes.NewReader([]byte(`{"email":"bagus@gmail.com", "password":"bagus123"}`))
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

func TestGetAllUsers(t *testing.T) {
	var mdl = usecase.MockUserUseCase{Status: "valid"}
	var conf = config.InitConfig()
	var ctl = NewUserController(&mdl, *conf)
	var e = echo.New()

	e.GET("/users", ctl.GetAllUsers(), echojwt.JWT([]byte(conf.Secret)))

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
	var conf = config.InitConfig()
	var mdl = usecase.MockUserUseCase{Status: "invalid"}
	var ctl = NewUserController(&mdl, *conf)
	var e = echo.New()

	e.POST("/users/login", ctl.LoginUser())

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
		var mdl = usecase.MockUserUseCase{Status: "ID < 1"}
		var conf = config.InitConfig()
		var ctl = NewUserController(&mdl, *conf)
		var e = echo.New()

		e.POST("/users/login", ctl.LoginUser())

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
		var mdl = usecase.MockUserUseCase{Status: "valid"}
		var conf = config.InitConfig()
		var ctl = NewUserController(&mdl, *conf)
		var e = echo.New()

		e.POST("/users/login", ctl.LoginUser())

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
