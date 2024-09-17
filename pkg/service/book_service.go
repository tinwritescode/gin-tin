package service

import (
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/repository"
	"gorm.io/gorm"
)

type BookService interface {
	GetAllBooks() ([]model.Book, error)
	CreateBook(book model.Book) error
	DeleteBook(id string) error
}

type bookService struct {
	repo repository.BookRepository
	db   *gorm.DB
}

func NewBookService(repo repository.BookRepository, db *gorm.DB) BookService {
	return &bookService{repo: repo, db: db}
}

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) CreateBook(book model.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}
