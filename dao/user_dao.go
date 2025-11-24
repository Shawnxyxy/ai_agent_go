package dao

import (
	"ai_agent/database"
	"ai_agent/model"
	"encoding/json"

	"context"
	"fmt"
	"time"
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

var ctx = context.Background()

// 获取用户时加缓存
func GetUserByIDWithCache(id uint) (*model.User, error) {
	key := fmt.Sprintf("user:%d", id)

	// 1. 尝试从 Redis 读取
	cached, err := database.RDB.Get(ctx, key).Result()
	if err == nil {
		// 缓存命中
		fmt.Println("Cache hit for user", id)
		var u model.User
		if err := json.Unmarshal([]byte(cached), &u); err == nil{
			return &u, nil
		}
	}
	// 2. 缓存未命中，查询数据库
	fmt.Println("Cache miss for user", id)
	var u model.User
	if err := database.DB.First(&u, id).Error; err != nil {
		return nil ,err
	}
	// 3. 写回缓存
	data, _ := json.Marshal(u)
	database.RDB.Set(ctx, key, data, 5*time.Minute)

	return &u, nil
}