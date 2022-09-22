package models

import (
	domains "agmc/day2/internal/domains"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

var books = []domains.Book{{Name: "lorem ipsum", Author: "John Doe", Year: 2022}, {Name: "lorem ipsum2", Author: "John Doe2", Year: 2022}}

func NewBookRepository(Conn *gorm.DB) domains.BookRepository {
	return &BookRepository{
		db: Conn,
	}
}

func (b *BookRepository) GetAllBooksRepository() *[]domains.Book {
	return &books
}

func (b *BookRepository) GetBookByIDRepository(id string) *domains.Book {
	return &books[0]
}

func (b *BookRepository) AddBookRepository(book *domains.Book) error {
	return nil
}

func (b *BookRepository) UpdateBookRepository(id string, book *domains.Book) error {
	return nil
}

func (b *BookRepository) DeleteBookRepository(id string) error {
	return nil
}
