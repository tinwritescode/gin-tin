package repository

import "github.com/tinwritescode/gin-tin/pkg/model"

type BookRepository interface {
	GetAll() []model.Book
	Create(book model.Book)
	Delete(id string)
}

type bookRepository struct {
	books []model.Book
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		books: []model.Book{
			{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
			{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
			{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
		},
	}
}

func (r *bookRepository) GetAll() []model.Book {
	return r.books
}

func (r *bookRepository) Create(book model.Book) {
	r.books = append(r.books, book)
}

func (r *bookRepository) Delete(id string) {
	for i, book := range r.books {
		if book.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			break
		}
	}
}
