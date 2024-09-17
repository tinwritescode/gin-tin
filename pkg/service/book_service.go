package service

import (
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/repository"
)

type BookService interface {
	GetAllBooks() ([]model.Book, error)
	CreateBook(book model.Book) (model.Book, error)
	DeleteBook(id string) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) CreateBook(book model.Book) (model.Book, error) {
	return s.repo.Create(book)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}
