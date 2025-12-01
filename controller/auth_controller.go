package controller

import (
	"ai_agent/service"
	"ai_agent/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 定义密钥
var jwtKey = []byte("my_secret_key")

func LoginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "invalid parameters")
		return
	}

	auth, err := service.LoginService(req.Username, req.Password)
	if err != nil {
		response.Fail(c, 401, err.Error())
		return
	}

	// 生成 JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": req.Username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })
    tokenString, _ := token.SignedString(jwtKey)

    response.Success(c, gin.H{
        "message": "login success",
        "token":   tokenString,
        "user":    auth,
    })
}