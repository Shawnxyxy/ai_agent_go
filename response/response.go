package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code	int			`json:"code"`
	Msg 	string		`json:"msg"`
	Data 	interface{} `json:"data,omitempty"`
}

func RespOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 	http.StatusOK,
		Msg: 	"OK",
		Data:	data,
	})
}

func RespError(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
	})
}