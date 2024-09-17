package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/service"
)

type Handler struct {
	bookService service.BookService
}

func NewHandler(bookService service.BookService) *Handler {
	return &Handler{bookService: bookService}
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	r.GET("/", h.getRoot)
	r.GET("/books", h.getBooks)
	r.POST("/books", h.createBook)
	r.DELETE("/books/:id", h.deleteBook)
}

func (h *Handler) getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func (h *Handler) getBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) createBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.bookService.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *Handler) deleteBook(c *gin.Context) {
	id := c.Param("id")
	err := h.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Implement handler methods
