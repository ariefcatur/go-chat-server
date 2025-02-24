package middleware

import (
	"github.com/gin-gonic/gin"
	"go-chat-server/helpers"
	"go-chat-server/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			helpers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{
				"details": "Missing token",
			})
			c.Abort()
			return
		}

		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			helpers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{
				"details": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
