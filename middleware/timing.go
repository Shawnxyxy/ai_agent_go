package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Timing() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // 执行下一个中间件
		cost := time.Since(start) // 计算整个请求耗时

		fmt.Printf("[Timing] %s %s took %v\n",
			c.Request.Method,
			c.Request.URL.Path,
			cost,
		)
	}
}