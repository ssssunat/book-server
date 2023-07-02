package service

import (
	"wc/model"
	"wc/pkg/repository"
)

type Book interface {
	GetBooksByAuthor(authorId int) ([]model.Book, error)
	GetAuthorBooksCount(authorId int) (int, error)
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos.Book),
	}
}
