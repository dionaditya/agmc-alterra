package domains

type Book struct {
	Name   string `gorm:"not null" json:"Name" validate:"required"`
	Author string `gorm:"not null" json:"Author" validate:"required"`
	Year   int    `gorm:"not null" json:"Year,omitempty" validate:"required"`
}

type BookRepository interface {
	GetAllBooksRepository() *[]Book
	GetBookByIDRepository(id string) *Book
	AddBookRepository(book *Book) error
	UpdateBookRepository(id string, book *Book) error
	DeleteBookRepository(id string) error
}

type BookUseCase interface {
	GetAllBooks() *[]Book
	GetBookByID(id string) *Book
	AddBook(book *Book) error
	UpdateBook(id string, book *Book) error
	DeleteBook(id string) error
}
