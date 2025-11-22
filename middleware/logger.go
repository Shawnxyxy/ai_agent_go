package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // 记录开始时间

		c.Next() // 执行下一个中间件或业务逻辑

		// 请求执行完后，统计信息
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		log.Printf("[INFO] %s %s %d %v", method, path, status, latency)
	}
}
