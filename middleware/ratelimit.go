package middleware

import (
	"ai_agent/response"
	"sync"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

// RateLimiter结构体
type RateLimiter struct {
	sync.Mutex
	Requests	int			// 当前计数
	LastReset	time.Time	// 上次重置时间
	Limit		int			// 每秒允许请求数
	IntervalSec int			// 统计周期，秒
}
// 创建限流器
func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		Requests:	0,
		LastReset:  time.Now(),
		Limit:		limit,
		IntervalSec: 1,
	}
}

func RateLimit(limiter *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter.Lock()
		defer limiter.Unlock()

		now := time.Now()
		if now.Sub(limiter.LastReset) >= time.Duration(limiter.IntervalSec)*time.Second{
			// 重置计数
			limiter.Requests = 0
			limiter.LastReset = now
		}
		if limiter.Requests >= limiter.Limit {
			response.Fail(c, http.StatusTooManyRequests, "too many requests")
			c.Abort()
			return
		}
		limiter.Requests++
		c.Next()
	}
}