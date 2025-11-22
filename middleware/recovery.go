package middleware

import (
	"log"
	"errors"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() { // 延迟执行函数
			if err := recover(); err != nil {
				log.Printf("panic:%v\nstack:%s", err, debug.Stack())
				c.Error(errors.New("Interval Server Error"))
				c.Abort() // 中断后续执行链
			}
		}()
		c.Next() // 继续执行后续中间件或路由逻辑
	}
}
