package controller

import (
    "ai_agent/response"
    "ai_agent/service"
    "net/http"

    "github.com/gin-gonic/gin"
)

type registerRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
    var req registerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Fail(c, http.StatusBadRequest, "invalid request body")
        return
    }

    if err := service.Register(req.Username, req.Password); err != nil {
        // 区分重复用户名与其他错误
        if err.Error() == "username already exists" {
            response.Fail(c, http.StatusBadRequest, err.Error())
            return
        }
        response.Fail(c, http.StatusInternalServerError, "register failed: "+err.Error())
        return
    }

    response.Success(c, gin.H{"message": "register success"})
}