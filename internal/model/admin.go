package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;size:100;not null"`
	PasswordHash string `gorm:"not null"`
}
