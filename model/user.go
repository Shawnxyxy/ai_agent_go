package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:50;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}