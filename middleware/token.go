package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Aungwaiyanchit/books/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var tokenString string
		barerToken := c.Request.Header.Get("Authorization")
		if len(strings.Split(barerToken, " ")) == 2 {
			tokenString =  strings.Split(barerToken, " ")[1]
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		var userClams models.JwtUser
		token, err := jwt.ParseWithClaims(tokenString, &userClams, func(token *jwt.Token) (interface{}, error) {
			return []byte("kitty harein"), nil
		})
		if err != nil {
			fmt.Println(err)
			fmt.Println(barerToken)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}