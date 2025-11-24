package dao

import (
	"ai_agent/database"
	"ai_agent/model"
)

func GetUserAuthByUsername(username string) (*model.UserAuth, error) {
	var ua model.UserAuth
	err := database.DB.Where("username = ?", username).First(&ua).Error
	return &ua, err
}