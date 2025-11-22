package controller

import (
	"ai_agent/response"
	"ai_agent/service"
	"ai_agent/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ====== CRUD Handlers ======

// 创建用户
func CreateUserHandler(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespError(c, http.StatusBadRequest, "invalid json body:"+err.Error())
		return
	}
	if err := service.CreateUserHandler(&req); err != nil {
		response.RespError(c, http.StatusInternalServerError, "db insert error: "+err.Error())
		return
	}
	response.RespOk(c, gin.H{
		"message": "user created successfully",
		"user": req,
	})
}
// 获取单个用户
func UserHandler(c *gin.Context) {
	idStr := c.Param("id") // 从路径中获取参数
	id ,err := strconv.Atoi(idStr)
	if err != nil {
		response.RespError(c, http.StatusBadRequest, "invalid user id")
		return
	}
	user, err := service.GetUserService(uint(id))
	if err != nil {
		response.RespError(c, http.StatusNotFound, "user not found")
		return
	}
	response.RespOk(c, user)
}
// 获取用户详情（支持 query 参数）
func UserDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.RespError(c, http.StatusBadRequest, "invalid user id")
		return
	}
	user, err := service.GetUserService(uint(id))
	if err != nil {
		response.RespError(c, http.StatusNotFound, "user not found")
		return
	}
	active := c.DefaultQuery("active", "false") // query 参数
	response.RespOk(c, gin.H{
		"user":     user,
		"is_active": active,
	})
}
// 更新用户
func UpdateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id ,err := strconv.Atoi(idStr)
	if err != nil {
		response.RespError(c, http.StatusBadRequest, "invalid user id")
		return
	}
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespError(c, http.StatusBadRequest, "invalid json body: "+err.Error())
		return
	}
	req.ID = uint(id)
	if err := service.UpdateUserService(&req); err != nil {
		response.RespError(c, http.StatusNotFound, "user not found")
		return
	}

	response.RespOk(c, gin.H{
		"message": "user updated successfully",
		"user":    req,
	})
}
// 删除用户
func DeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.RespError(c, http.StatusBadRequest, "invalid user id")
		return
	}
	if err := service.DeleteUserService(uint(id)); err != nil {
		response.RespError(c, http.StatusInternalServerError, "db delete error: "+err.Error())
		return
	}
	response.RespOk(c, gin.H{
		"message": "user deleted successfully",
	})
}