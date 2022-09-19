package usecase

import (
	domains "agmc/day2/internal/domains"
)

type BookUseCase struct {
	bookRepository domains.BookRepository
}

func NewBookUseCase(bookRepo domains.BookRepository) domains.BookUseCase {
	return &BookUseCase{
		bookRepository: bookRepo,
	}
}

func (b *BookUseCase) GetAllBooks() *[]domains.Book {
	return b.bookRepository.GetAllBooksRepository()
}

func (b *BookUseCase) GetBookByID(id string) *domains.Book {
	return b.bookRepository.GetBookByIDRepository(id)
}

func (b *BookUseCase) AddBook(book *domains.Book) error {
	return b.bookRepository.AddBookRepository(book)
}

func (b *BookUseCase) UpdateBook(id string, book *domains.Book) error {
	return b.bookRepository.UpdateBookRepository(id, book)
}

func (b *BookUseCase) DeleteBook(id string) error {
	return b.bookRepository.DeleteBookRepository(id)
}
