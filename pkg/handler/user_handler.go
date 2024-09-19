package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "List of users",
		"users":   users,
	})
}
