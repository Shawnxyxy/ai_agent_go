package router

import (
	"ai_agent/controller"
	"ai_agent/middleware"
	"ai_agent/response"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Gin 初始化了一个 HTTP 服务器引擎对象 r
	// Default方法自带默认中间件（日志与恢复）
	// New方法没有任何中间件
	r := gin.Default()

	// 注册自定义的中间件
	r.Use(middleware.Logger())
	r.Use(middleware.CustomRecovery())
	r.Use(middleware.Timing())
	limiter := middleware.NewRateLimiter(5)
	r.Use(middleware.RateLimit(limiter))
	r.Use(middleware.Auth())
	r.Use(middleware.ErrorHandler())

	// ====== 系统级接口 ======
	// health接口验证
	r.GET("/health", func(c *gin.Context) {
		response.RespOk(c, gin.H{"status": "ok"})
	})
	// recovery中间件验证
	r.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})
	// =======================

	// ====== 业务路由分组 ======
	api := r.Group("/api") // 创建子路由组，统一前缀/api
	{
		api.GET("/ping", controller.PingHandler)     // 在组下注册 GET 路由，实际路径 /api/ping
		api.GET("/user/:id", controller.UserHandler) // /:id 的意思是会匹配 URL 中这一段的值，并自动传给你
		api.GET("/user/:id/detail", controller.UserDetailHandler)
		api.POST("/user/create", controller.CreateUserHandler)
		api.PUT("/user/:id", controller.UpdateUserHandler)
		api.DELETE("/user/:id", controller.DeleteUserHandler)
	}
	// ========================
	return r
}
