package model

import "gorm.io/gorm"

type UserAuth struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}