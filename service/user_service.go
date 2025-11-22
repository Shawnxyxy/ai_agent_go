package service

import (
	"ai_agent/dao"
	"ai_agent/model"
	"errors"
)

// 创建用户
func CreateUserHandler(u *model.User) error {
	if u.Name == "" || u.Email == "" {
		return errors.New("name or email cannot be empty")
	}
	return dao.InserUser(u)
}
// 根据ID获取用户
func GetUserService(id uint) (*model.User, error) {
	return dao.GetUserByID(id)
}
// 更新用户
func UpdateUserService(u *model.User) error {
	_, err := dao.GetUserByID(u.ID)
	if err != nil {
		return err
	}
	return dao.UpdateUser(u)
}
// 删除用户
func DeleteUserService(id uint) error {
	return dao.DeleteUser(id)
}