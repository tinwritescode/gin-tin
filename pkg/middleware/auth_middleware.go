package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tinwritescode/gin-tin/pkg/config"
	"github.com/tinwritescode/gin-tin/pkg/model"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "Authorization header is required",
			})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "Invalid token format",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "Invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "Invalid token claims",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("user_role", claims["role"]) // Add this line to set the user's role
		c.Next()
	}
}
