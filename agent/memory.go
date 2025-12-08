package agent

import (
	"time"
	"ai_agent/database"
	"ai_agent/model"
	"ai_agent/llm"
)

func SaveMemory(userID uint, content string) error {
	memoryType := classifyMemory(content)
	importance := scoreImportance(content)
	mem := model.Memory{
		UserID:  	userID,
		Content: 	content,
		Type:   	memoryType,
		Importance: importance,
	}
	return database.DB.Create(&mem).Error
}

func LoadRecentMemory(userID uint, limit int) ([]model.Memory, error) {
	var mems []model.Memory
	err := database.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&mems).Error
	
	return mems, err
}

func buildMemoryContext(mems []model.Memory) string {
	if len(mems) == 0 {
		return "No prior memory."
	}
	result := "Recent memory:\n"
	for _, m := range mems {
		result += "-" + m.Content + "\n"
 	}
	return result
}

// 清理低价值记忆
func RecycleMemory(userID uint) error {
	// 1. 删除 Importance=0 且 7 天前的记忆
	threshold := time.Now().AddDate(0, 0, -7) // 7 天前
	err := database.DB.
		Where("user_id = ? AND importance = 0 AND created_at < ?", userID, threshold).
		Delete(&model.Memory{}).Error
	if err != nil {
		return err
	}
	return nil
}

// 压缩最近低重要度记忆
func CompressMemory(userID uint) error {
	var mems []model.Memory
	err := database.DB.
		Where("user_id = ? AND importance = 0", userID).
		Order("created_at DESC").
		Find(&mems).Error
	if err != nil {
		return err
	}
	// 合并成一句话
	combined := ""
	for _, m := range mems {
		combined += m.Content + " | "
	}
	// 让 LLM 压缩成一句话
	summary, _ := llm.AskDeepSeek("请将以下内容压缩成一句话记忆：" + combined)
	// 保存压缩后的记忆
	if summary != "" {
		SaveMemory(userID, summary)
	}
	// 删除原有低重要度记忆
	for _, m := range mems {
		database.DB.Delete(&m)
	}
	return nil
}