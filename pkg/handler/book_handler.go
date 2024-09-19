package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/utils"
)

func (h *Handler) getBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) createBook(c *gin.Context) {
	var bookRequest struct {
		Title       string `json:"title" validate:"required"`
		Author      string `json:"author" validate:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(bookRequest); err != nil {
		utils.HandleValidationErrors(c, err)
		return
	}

	userID := c.GetUint("user_id")

	book := model.Book{
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
		UserID:      userID,
	}

	createdBook, err := h.bookService.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdBook)
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
