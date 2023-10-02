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

type BlogMock struct {
	status string
}

func (blogMock *BlogMock) GetAllBlogs() []model.Blog {
	if blogMock.status == "valid" {
		listBlog := []model.Blog{
			{Judul: "Buku 1"}, {Judul: "Buku 2"}, {Judul: "Buku 3"},
		}
		return listBlog
	}
	return nil
}

func (blogMock *BlogMock) GetBlogById(id int) model.Blog {
	if blogMock.status == "valid" {
		result := model.Blog{
			Model:  gorm.Model{ID: 1},
			UserID: 1,
			Judul:  "Hello",
			Konten: "Hello World",
		}
		return result
	}
	return model.Blog{}
}

func (blogMock *BlogMock) CreateBlog(newBlog model.Blog) *model.Blog {
	if blogMock.status == "valid" {
		return &newBlog
	}
	return nil
}

func (blogMock *BlogMock) UpdateBlog(updateBlog model.Blog) *model.Blog {
	if blogMock.status == "valid" {
		return &updateBlog
	}
	return nil
}

func (blogMock *BlogMock) DeleteBlog(id int) {}

func TestGetAllBlogs(t *testing.T) {
	var mdl = BlogMock{status: "valid"}
	var config = configs.InitConfig()
	var ctl = NewBlogController(&mdl)
	var e = echo.New()

	e.GET("/blogs", ctl.GetBlogs, echojwt.JWT([]byte(config.Secret)))

	var req = httptest.NewRequest(http.MethodGet, "/blogs", nil)
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
	assert.Equal(t, "success get all blogs", tmp.Message)
}

func TestGetBlogById(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		var mdl = BlogMock{status: "invalid"}
		var config = configs.InitConfig()
		var ctl = NewBlogController(&mdl)
		var e = echo.New()

		e.GET("/blogs/:id", ctl.GetBlog, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodGet, "/blogs/satu", nil)
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
		var mdl = BlogMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewBlogController(&mdl)
		var e = echo.New()

		e.GET("/blogs/:id", ctl.GetBlog, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodGet, "/blogs/1", nil)
		req.Header.Set(echo.HeaderAuthorization, Token)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success get blog", tmp.Message)
	})
}

func TestCreateBlog(t *testing.T) {
	var mdl = BlogMock{status: "invalid"}
	var ctl = NewBlogController(&mdl)
	var e = echo.New()

	e.POST("/blogs", ctl.Create)

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"user_id":1, "judul":"Hello", "konten":"Hello World"}`))
		var req = httptest.NewRequest(http.MethodPost, "/blogs", bodyReq)
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
		var bodyReq = bytes.NewReader([]byte(`{"user_id":1, "judul":"Hello", "konten":"Hello World"}`))
		var req = httptest.NewRequest(http.MethodPost, "/blogs", bodyReq)
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
		var mdl = BlogMock{status: "valid"}
		var ctl = NewBlogController(&mdl)
		var e = echo.New()

		e.POST("/blogs", ctl.Create)

		var bodyReq = bytes.NewReader([]byte(`{"user_id":1, "judul":"Hello", "konten":"Hello World"}`))
		var req = httptest.NewRequest(http.MethodPost, "/blogs", bodyReq)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NotNil(t, tmp.Data)
		assert.Equal(t, "success create blog", tmp.Message)
	})
}

func TestUpdateBlog(t *testing.T) {
	var mdl = BlogMock{status: "invalid"}
	var config = configs.InitConfig()
	var ctl = NewBlogController(&mdl)
	var e = echo.New()

	e.PUT("/blogs/:id", ctl.Update, echojwt.JWT([]byte(config.Secret)))

	t.Run("Bind error", func(t *testing.T) {
		var bodyReq = bytes.NewReader([]byte(`{"user_id":1, "judul":"Hello", "konten":"Hello World"}`))
		var req = httptest.NewRequest(http.MethodPut, "/blogs/1", bodyReq)
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
		var bodyReq = bytes.NewReader([]byte(`{"user_id":1, "judul":"Hello", "konten":"Hello World"}`))
		var req = httptest.NewRequest(http.MethodPut, "/blogs/satu", bodyReq)
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
		var mdl = BlogMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewBlogController(&mdl)
		var e = echo.New()

		e.PUT("/blogs/:id", ctl.Update, echojwt.JWT([]byte(config.Secret)))

		var bodyReq = bytes.NewReader([]byte(`{"user_id":1, "judul":"Hello", "konten":"Hello World"}`))
		var req = httptest.NewRequest(http.MethodPut, "/blogs/1", bodyReq)
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

func TestDeleteBlog(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		var mdl = BlogMock{status: "invalid"}
		var config = configs.InitConfig()
		var ctl = NewBlogController(&mdl)
		var e = echo.New()

		e.DELETE("/blogs/:id", ctl.Delete, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodDelete, "/blogs/satu", nil)
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
		var mdl = BlogMock{status: "valid"}
		var config = configs.InitConfig()
		var ctl = NewBlogController(&mdl)
		var e = echo.New()

		e.DELETE("/blogs/:id", ctl.Delete, echojwt.JWT([]byte(config.Secret)))

		var req = httptest.NewRequest(http.MethodDelete, "/blogs/1", nil)
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
