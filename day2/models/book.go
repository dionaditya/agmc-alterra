package models

import (
	"agmc/day2/config"
	"agmc/day2/entity"
)

type BookModel struct {
	db config.DatabaseConfig
}

var books = []entity.Book{entity.Book{Name: "lorem ipsum", Author: "John Doe", Year: 2022}, entity.Book{Name: "lorem ipsum2", Author: "John Doe2", Year: 2022}}

func (e *BookModel) GetAllBooks() *[]entity.Book {
	return &books
}

func (e *BookModel) GetBookByID(id string) *entity.Book {
	return &books[0]
}

func (e *BookModel) AddBook(book *entity.Book) error {
	return nil
}

func (e *BookModel) UpdateBook(id string, book *entity.Book) error {
	return nil
}

func (e *BookModel) DeleteBook(id string) error {
	return nil
}
