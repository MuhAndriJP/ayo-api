package model

import "gorm.io/gorm"

type Position string

const (
	PositionPenyerang     Position = "penyerang"
	PositionGelandang     Position = "gelandang"
	PositionBertahan      Position = "bertahan"
	PositionPenjagaGawang Position = "penjaga_gawang"
)

type Player struct {
	gorm.Model
	TeamID       int64    `gorm:"not null;index"`
	Team         Team     `gorm:"constraint:OnDelete:RESTRICT"`
	Name         string   `gorm:"size:200;not null"`
	HeightCm     float32  `gorm:"not null"`
	WeightKg     float32  `gorm:"not null"`
	Position     Position `gorm:"type:enum('penyerang','gelandang','bertahan','penjaga_gawang');not null"`
	JerseyNumber int64    `gorm:"not null"`
}

func PositionFromString(s string) Position {
	return Position(s)
}

func (p Position) String() string {
	return string(p)
}
