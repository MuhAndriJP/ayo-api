package model

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	MatchID      int64  `gorm:"not null;index"`
	Match        Match  `gorm:"constraint:OnDelete:CASCADE"`
	PlayerID     int64  `gorm:"not null;index"`
	Player       Player `gorm:"constraint:OnDelete:RESTRICT"`
	MinuteScored int64  `gorm:"not null"`
}
