package factory

import (
	bookRepo "agmc/day2/internal/book/repository"
	"agmc/day2/internal/domains"
	userRepo "agmc/day2/internal/user/repository"

	"gorm.io/gorm"
)

type Factory struct {
	UserRepo domains.UserRepository
	BookRepo domains.BookRepository
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		UserRepo: userRepo.NewUserRepositoryy(db),
		BookRepo: bookRepo.NewBookRepository(db),
	}
}
