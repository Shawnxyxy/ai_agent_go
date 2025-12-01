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