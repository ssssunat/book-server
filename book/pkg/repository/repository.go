package repository

import (
	"wc/model"
	"github.com/jmoiron/sqlx"
)

type Book interface {
	GetBooksByAuthor(authorId int) ([]model.Book, error)
	GetAuthorBooksCount(authorId int) (int, error)
}

type Repository struct {
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Book: NewBookPostgres(db),
	}
}
