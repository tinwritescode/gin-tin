package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinwritescode/gin-tin/pkg/config"
	"github.com/tinwritescode/gin-tin/pkg/middleware"
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/service"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	bookService service.BookService
	authService service.AuthService
	userService service.UserService
	config      *config.Config
}

func NewHandler(bookService service.BookService, authService service.AuthService, userService service.UserService, config *config.Config) *Handler {
	return &Handler{bookService: bookService, authService: authService, userService: userService, config: config}
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	})

	// Public routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/refresh", h.RefreshToken)
	r.POST("/logout", h.Logout)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(h.config))
	{
		protected.GET("/books", h.getBooks)
		protected.POST("/books", h.createBook)
		protected.DELETE("/books/:id", h.deleteBook)
	}

	r.GET("/", h.getRoot)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Admin routes
	admin := protected.Group("/")
	admin.Use(middleware.RoleMiddleware(model.RoleAdmin, model.RoleSuperAdmin))
	{
		// Admin specific routes
		admin.GET("/users", h.getUsers) // This line remains the same
	}
}

func (h *Handler) getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
