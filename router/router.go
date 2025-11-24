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

	// ====== 公共路由分组（无需JWT） ======
	public := r.Group("/api") // 创建子路由组，统一前缀/api
	{
		public.POST("/register", controller.RegisterHandler)
		public.POST("/login", controller.LoginHandler)
		public.GET("/ping", controller.PingHandler)
	}
	// ====== 私有接口（需要JWT） ======
	private := r.Group("/api")
	private.Use(middleware.JWTAuthMiddleware()) // 只对私有接口加JWT
	{
		private.GET("/user/:id", controller.UserHandler) // 在组下注册 GET 路由，实际路径 /api/ping
		private.GET("/user/:id/detail", controller.UserDetailHandler) // /:id 的意思是会匹配 URL 中这一段的值，并自动传给你
		private.POST("/user/create", controller.CreateUserHandler)
		private.PUT("/user/:id", controller.UpdateUserHandler)
		private.DELETE("/user/:id", controller.DeleteUserHandler)
	}

	// ========================
	return r
}
