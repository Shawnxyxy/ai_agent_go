package dao

import (
	"ai_agent/database"
	"ai_agent/model"
)
// 插入用户
func InserUser(u *model.User) error {
	return database.DB.Create(u).Error
}
// 根据ID获取用户
func GetUserByID(id uint) (*model.User, error) {
	var u model.User
	err := database.DB.First(&u, id).Error
	return &u, err
}
// 更新用户
func UpdateUser(u *model.User) error {
	return database.DB.Save(u).Error
}
// 删除用户
func DeleteUser(id uint) error {
	return database.DB.Delete(&model.User{}, id).Error
}