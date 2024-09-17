package repository

import (
	"errors"
	"strings"

	"github.com/tinwritescode/gin-tin/pkg/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]model.Book, error)
	Create(book model.Book) (model.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book

	if result := r.db.Find(&books); result.Error != nil {
		return nil, result.Error
	}

	return books, nil

}

func (r *bookRepository) Create(book model.Book) (model.Book, error) {
	result := r.db.Create(&book)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			return model.Book{}, errors.New("book already exists")
		}
		return model.Book{}, result.Error
	}
	createdBook := *result.Statement.Dest.(*model.Book)

	return createdBook, nil
}

func (r *bookRepository) Delete(id string) error {
	if result := r.db.Delete(&model.Book{}, id); result.Error != nil {
		return result.Error
	}

	return nil
}
