package service

import (
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/repository"
)

type BookService interface {
	GetAllBooks() []model.Book
	CreateBook(book model.Book)
	DeleteBook(id string)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() []model.Book {
	return s.repo.GetAll()
}

func (s *bookService) CreateBook(book model.Book) {
	s.repo.Create(book)
}

func (s *bookService) DeleteBook(id string) {
	s.repo.Delete(id)
}
