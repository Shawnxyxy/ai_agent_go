package model

import "time"

type Task struct {
	ID			uint      `gorm:"primaryKey" json:"id"`
	Type    	string    `json:"type"`
	Payload		string    `json:"payload"`
	Status      string    `json:"status"`
	Result 		string    `json:"result"`
	RetryCount	int       `json:"retryCount"`
	CreatedAt	time.Time `json:"createdAt"`
	UpdatedAt	time.Time `json:"updatedAt"`
}

const (
	TaskTypeEcho       = "echo"
	TaskTypeInitMemory = "init_memory"
	TaskTypeWarmup     = "warmup"
	TaskTypeLLMCall    = "llm_call"
)

type Memory struct {
    ID         uint      `gorm:"primaryKey"`
    UserID     uint      `gorm:"default:0"`
    Content    string    `gorm:"type:text"`
    Type       string    `gorm:"type:varchar(20);default:'interaction'"` // memory type
    Importance int       `gorm:"default:0"`                              // importance score
    CreatedAt  time.Time `gorm:"autoCreateTime"`
}