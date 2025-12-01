package controller

import (
	"ai_agent/response"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) { // gin.Contex类型的指针
	aStr := c.DefaultQuery("a", "10")
	bStr, exists := c.GetQuery("b")
	if !exists {
		response.Fail(c, http.StatusBadRequest, "missing parameter b")
		return
	}
	a, _ := strconv.Atoi(aStr) // 把字符串 aStr、bStr 转换成整数，分别存入变量 a 和 b。
	b, _ := strconv.Atoi(bStr) // 不需要 err，所以用 _ 来忽略
	// 返回 JSON 响应
	response.Success(c, gin.H{
		"result": a + b,
	})
}
