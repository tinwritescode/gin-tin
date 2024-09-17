package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinwritescode/gin-tin/pkg/middleware"
	"github.com/tinwritescode/gin-tin/pkg/service"
)

type Handler struct {
	bookService service.BookService
	authService service.AuthService
}

func NewHandler(bookService service.BookService, authService service.AuthService) *Handler {
	return &Handler{bookService: bookService, authService: authService}
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/register", h.register)
	r.POST("/login", h.login)
	r.POST("/refresh", h.refreshToken)
	r.POST("/logout", h.logout)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/books", h.getBooks)
		protected.POST("/books", h.createBook)
		protected.DELETE("/books/:id", h.deleteBook)
	}

	r.GET("/", h.getRoot)
}

func (h *Handler) getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
