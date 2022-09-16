package controllers

import (
	"agmc/day2/entity"
	"agmc/day2/models"
	"agmc/day2/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookModel models.BookModel
}

func (BookController *BookController) GetAllBooks(c echo.Context) error {
	books := BookController.bookModel.GetAllBooks()
	return utils.HandleSuccess(c, books)
}

func (BookController *BookController) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	book := BookController.bookModel.GetBookByID(id)
	return utils.HandleSuccess(c, book)
}

func (BookController *BookController) AddBook(c echo.Context) error {
	var book = entity.Book{}

	c.Bind(&book)

	bookData := BookController.bookModel.AddBook(&book)

	return utils.HandleSuccess(c, bookData)
}

func (BookController *BookController) UpdateBook(c echo.Context) error {
	var book = entity.Book{}
	id := c.Param("id")

	c.Bind(&book)

	bookData := BookController.bookModel.UpdateBook(id, &book)

	return utils.HandleSuccess(c, bookData)
}

func (BookController *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")

	err := BookController.bookModel.DeleteBook(id)

	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.HandleSuccess(c, struct{}{})
}
