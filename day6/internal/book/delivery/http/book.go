package controllers

import (
	"agmc/day2/internal/book/usecase"
	"agmc/day2/internal/domains"
	"agmc/day2/internal/factory"
	"agmc/day2/pkg/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookUseCase domains.BookUseCase
}

func NewBookHanndler(f *factory.Factory) BookHandler {
	return BookHandler{
		bookUseCase: usecase.NewBookUseCase(f.BookRepo),
	}
}

func (b *BookHandler) GetAllBooks(c echo.Context) error {
	books := b.bookUseCase.GetAllBooks()
	return response.HandleSuccess(c, books)
}

func (b *BookHandler) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	book := b.bookUseCase.GetBookByID(id)
	return response.HandleSuccess(c, book)
}

func (b BookHandler) AddBook(c echo.Context) error {
	var book = domains.Book{}

	c.Bind(&book)

	bookData := b.bookUseCase.AddBook(&book)

	return response.HandleSuccess(c, bookData)
}

func (b *BookHandler) UpdateBook(c echo.Context) error {
	var book = domains.Book{}
	id := c.Param("id")

	c.Bind(&book)

	bookData := b.bookUseCase.UpdateBook(id, &book)

	return response.HandleSuccess(c, bookData)
}

func (b *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")

	err := b.bookUseCase.DeleteBook(id)

	if err != nil {
		return response.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	return response.HandleSuccess(c, struct{}{})
}
