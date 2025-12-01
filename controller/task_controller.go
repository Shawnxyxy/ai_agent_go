package controller

import (
	"ai_agent/response"
	"ai_agent/service"
	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func CreateTaskHandler(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "Invalid request")
		return
	}

	task, err := service.CreateTask(req.Type, req.Payload)
	if err != nil {
		response.Fail(c, 500, "Failed to create task")
		return
	}

	response.Success(c, task)
}