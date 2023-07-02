package service

import (
	"wc/model"
	"wc/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooksByAuthor(authorId int) ([]model.Book, error) {
	return s.repo.GetBooksByAuthor(authorId)
}

func (s *BookService) GetAuthorBooksCount(authorId int) (int, error) {
	return s.repo.GetAuthorBooksCount(authorId)
}
