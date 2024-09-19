package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/utils"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "User registration details"
// @Success 201 {object} model.RegisterResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleValidationErrors(c, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		utils.HandleValidationErrors(c, err)
		return
	}

	err := h.authService.Register(user)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.username") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return access token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.LoginRequest true "User login credentials"
// @Success 200 {object} model.LoginResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleValidationErrors(c, err)
		return
	}

	accessToken, refreshToken, err := h.authService.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set refresh token as an HTTP-only cookie
	c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body model.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} model.RefreshTokenResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /refresh [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	newAccessToken, newRefreshToken, err := h.authService.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set new refresh token as an HTTP-only cookie
	c.SetCookie("refresh_token", newRefreshToken, 7*24*60*60, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

// Logout godoc
// @Summary Logout user
// @Description Invalidate user's access token
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.LogoutResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /logout [post]
func (h *Handler) Logout(c *gin.Context) {
	// Clear the refresh token cookie
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
