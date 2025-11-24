package middleware

import (
	"github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "net/http"
)

// 全局密钥
var jwtKey = []byte("my_secret_key")

func JWTAuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "missing token"})
            c.Abort()
            return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
		})
		if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "invalid token"})
            c.Abort()
            return
        }

        c.Next()
	}
}