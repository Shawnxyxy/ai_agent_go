package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`           // 0 表示成功（推荐规范）
	Msg  string      `json:"msg"`            // 描述信息
	Data interface{} `json:"data,omitempty"` // omitempty 为空时不返回
}

// 通用成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// 带自定义消息的成功
func SuccessMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

// 通用错误（业务错误）
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
}

// 系统级错误（不建议暴露详细原因）
func ServerError(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  "internal server error",
	})
}