package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name        string `gorm:"size:200;not null"`
	LogoPath    string `gorm:"size:500"`
	FoundedYear int64  `gorm:"not null"`
	Address     string `gorm:"size:500"`
	City        string `gorm:"size:100;not null"`
	Players     []Player
}
