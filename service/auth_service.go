package service

import (
	"errors"
    "strings"

    "ai_agent/database"
    "ai_agent/model"
    "ai_agent/utils"
	"ai_agent/dao"

	"golang.org/x/crypto/bcrypt"
)
// Register 注册新用户（保存 username + bcrypt(password) 到 user_auths 表）
func Register(username, password string) error {
	username = strings.TrimSpace(username)
	if username == "" || password == "" {
		return errors.New("username and password cannot be empty")
	}

	var existing model.UserAuth
	if err := database.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		return errors.New("username already exists")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	auth := model.UserAuth{
		Username: username,
		Password: hash,
	}
	if err := database.DB.Create(&auth).Error; err != nil {
		return err
	}
	return nil
}

func LoginService(username, password string) (*model.UserAuth, error) {
	auth, err := dao.GetUserAuthByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}
	return auth, nil
}