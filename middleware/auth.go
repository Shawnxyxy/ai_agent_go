package middleware

import (
	"ai_agent/response"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Auth 中间件：检查请求 Header 中是否包含 token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(("token"))
		if token == "" {
			response.RespError(c, http.StatusUnauthorized, "missing token")
			c.Abort()
			return
		}
		// token 存在，继续执行后续中间件和处理函数
		c.Next()
	}
}