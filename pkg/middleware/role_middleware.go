package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		fmt.Println(userRole)

		userRoleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user role format"})
			c.Abort()
			return
		}

		fmt.Println(userRoleStr)

		for _, role := range roles {
			if userRoleStr == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
	}
}
