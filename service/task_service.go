package service

import (
	"ai_agent/database"
	"ai_agent/model"
)

func CreateTask(taskType, payload string) (*model.Task, error) {
	task := &model.Task{
		Type:   taskType,
		Payload: payload,
		Status: "pending",
	}
	err := database.DB.Create(&task).Error
	return task, err
}

func GetPendingTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := database.DB.Where("status = ?", "pending").Find(&tasks).Error
	return tasks, err
}

func UpdateTaskStatus(taskID uint, status, result string) error {
	return database.DB.Model(&model.Task{}).
		Where("id = ?", taskID).
		Updates(map[string]interface{}{
			"status": status,
			"result": result,
		}).Error
}