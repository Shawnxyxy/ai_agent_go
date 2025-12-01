package middleware

import (
	"ai_agent/response"
	"log"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, ginErr := range c.Errors {
				log.Printf("[ERROR] %v\n", ginErr.Err)
			}
			response.Fail(c, 500, "Internal Server Error")
			c.Abort()
		}
	}
}