package controller

import (
	"REST_API_testing/helper"
	"REST_API_testing/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookControllerInterface interface {
	GetBooks(c echo.Context) error
	GetBook(c echo.Context) error
	Insert(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type BookController struct {
	model model.BookModelInterface
}

func NewBookController(bookModelInterface model.BookModelInterface) BookControllerInterface {
	return &BookController{
		model: bookModelInterface,
	}
}

// func (bookController *BookController) InitBookController(bookModel model.BookModel) {
// 	bookController.model = bookModel
// }

func (bookController *BookController) GetBooks(c echo.Context) error {
	var result = bookController.model.GetAllBooks()

	return c.JSON(http.StatusOK, helper.SetResponse("success get all books", result))
}

func (bookController *BookController) GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse("Invalid ID", nil))
	}

	var result = bookController.model.GetBookById(id)

	return c.JSON(http.StatusOK, helper.SetResponse("success get book", result))
}

func (bookController *BookController) Insert(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	var result = bookController.model.InsertBook(book)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
	}

	return c.JSON(http.StatusCreated, helper.SetResponse("success insert book", result))
}

func (bookController *BookController) Update(c echo.Context) error {
	var update model.Book
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse("Invalid ID", nil))
	}

	update.ID = uint(id)
	var result = bookController.model.UpdateBook(update)

	return c.JSON(http.StatusOK, helper.SetResponse("success update", result))
}

func (bookController *BookController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse("Invalid ID", nil))
	}

	bookController.model.DeleteBook(id)

	return c.JSON(http.StatusOK, nil)
}
