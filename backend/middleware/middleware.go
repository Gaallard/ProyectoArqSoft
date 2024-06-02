package middleware

import (
	client "backend/clients/users"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idUsuario := c.GetHeader("Authorization")
		if idUsuario == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No UserId provided"})
			c.Abort()
			return
		}
		id, _ := strconv.Atoi(strings.Replace(idUsuario, "Bearer ", "", 1))
		user, _ := client.GetUserById(id)
		if user.Role != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "usuario no autorizado"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
