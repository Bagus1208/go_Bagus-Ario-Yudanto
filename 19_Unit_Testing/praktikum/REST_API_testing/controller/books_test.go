package controller

import (
	"REST_API_testing/configs"
	"REST_API_testing/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type BookMock struct {
	status string
}

func (bookMock *BookMock) GetAllBooks() []model.Book {
	if bookMock.status == "valid" {
		listBook := []model.Book{
			{Judul: "Buku 1"}, {Judul: "Buku 2"}, {Judul: "Buku 3"},
		}
		return listBook
	}
	return nil
}

func (bookMock *BookMock) GetBookById(id int) model.Book {
	if bookMock.status == "valid" {
		result := model.Book{
			Model:    gorm.Model{ID: 1},
			Judul:    "Buku 1",
			Penulis:  "Rudy",
			Penerbit: "Gramedia",
		}
		return result
	}
	return model.Book{}
}

func (bookMock *BookMock) InsertBook(newBook model.Book) *model.Book {
	if bookMock.status == "valid" {
		return &newBook
	}
	return nil
}

func (bookMock *BookMock) UpdateBook(updateBook model.Book) *model.Book {
	if bookMock.status == "valid" {
		return &updateBook
	}
	return nil
}

func (bookMock *BookMock) DeleteBook(id int) {}

func TestGetAllBooks(t *testing.T) {
	var mdl = BookMock{status: "valid"}
	var config = configs.InitConfig()
	var ctl = NewBookController(&mdl)
	var e = echo.New()

	e.GET("/books", ctl.GetBooks, echojwt.JWT([]byte(config.Secret)))

	var req = httptest.NewRequest(http.MethodGet, "/books", nil)
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
	assert.Equal(t, "success get all books", tmp.Message)
}

func TestGetBookById(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		var mdl = BookMock{status: "invalid"}
		var config = configs.InitConfig()
		var ctl = NewBookController(&mdl)
		var e = echo.New()

		e.GET("/books/:id", ctl.GetBook, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodGet, "/books/satu", nil)
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
		var mdl = BookMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewBookController(&mdl)
		var e = echo.New()

		e.GET("/books/:id", ctl.GetBook, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodGet, "/books/1", nil)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success get book", tmp.Message)
	})
}

func TestInsertBook(t *testing.T) {
	var mdl = BookMock{status: "invalid"}
	var ctl = NewBookController(&mdl)
	var e = echo.New()

	e.POST("/books", ctl.Insert)

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"judul":"Buku 1", "Penulis":"Rudy", "Penerbit":"Gramedia"}`))
		var req = httptest.NewRequest(http.MethodPost, "/books", bodyReq)
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
		var bodyReq = bytes.NewReader([]byte(`{"judul":"Buku 1", "Penulis":"Rudy", "Penerbit":"Gramedia"}`))
		var req = httptest.NewRequest(http.MethodPost, "/books", bodyReq)
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
		var mdl = BookMock{status: "valid"}
		var ctl = NewBookController(&mdl)
		var e = echo.New()

		e.POST("/books", ctl.Insert)

		var bodyReq = bytes.NewReader([]byte(`{"judul":"Buku 1", "Penulis":"Rudy", "Penerbit":"Gramedia"}`))
		var req = httptest.NewRequest(http.MethodPost, "/books", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success insert book", tmp.Message)
	})
}

func TestUpdateBook(t *testing.T) {
	var mdl = BookMock{status: "invalid"}
	var config = configs.InitConfig()
	var ctl = NewBookController(&mdl)
	var e = echo.New()

	e.PUT("/books/:id", ctl.Update, echojwt.JWT([]byte(config.Secret)))

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"judul":"Buku 1", "Penulis":"Rudy", "Penerbit":"Gramedia"}`))
		var req = httptest.NewRequest(http.MethodPut, "/books/1", bodyReq)
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
		var req = httptest.NewRequest(http.MethodPut, "/books/satu", bodyReq)
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
		var mdl = BookMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewBookController(&mdl)
		var e = echo.New()

		e.PUT("/books/:id", ctl.Update, echojwt.JWT([]byte(config.Secret)))

		var bodyReq = bytes.NewReader([]byte(`{"judul":"Buku 1", "Penulis":"Rudy", "Penerbit":"Gramedia"}`))
		var req = httptest.NewRequest(http.MethodPut, "/books/1", bodyReq)
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

func TestDeleteBook(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		var mdl = BookMock{status: "invalid"}
		var config = configs.InitConfig()
		var ctl = NewBookController(&mdl)
		var e = echo.New()

		e.DELETE("/books/:id", ctl.Delete, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodDelete, "/books/satu", nil)
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
		var mdl = BookMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewBookController(&mdl)
		var e = echo.New()

		e.DELETE("/books/:id", ctl.Delete, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodDelete, "/books/1", nil)
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
