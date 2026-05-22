package model

import (
	"time"

	"gorm.io/gorm"
)

type MatchStatus string

const (
	MatchStatusScheduled MatchStatus = "scheduled"
	MatchStatusFinished  MatchStatus = "finished"
)

const (
	ResultHomeWin = "Tim Home Menang"
	ResultAwayWin = "Tim Away Menang"
	ResultDraw    = "Draw"
)

type Match struct {
	gorm.Model
	MatchDate  time.Time   `gorm:"type:date;not null"`
	MatchTime  string      `gorm:"type:varchar(8);not null"`
	HomeTeamID int64       `gorm:"not null;index"`
	HomeTeam   Team        `gorm:"foreignKey:HomeTeamID;constraint:OnDelete:RESTRICT"`
	AwayTeamID int64       `gorm:"not null;index"`
	AwayTeam   Team        `gorm:"foreignKey:AwayTeamID;constraint:OnDelete:RESTRICT"`
	HomeScore  int64       `gorm:"default:0"`
	AwayScore  int64       `gorm:"default:0"`
	Status     MatchStatus `gorm:"type:enum('scheduled','finished');default:'scheduled';not null"`
	Goals      []Goal
}

func MatchStatusFromString(s string) MatchStatus {
	return MatchStatus(s)
}

func (m MatchStatus) String() string {
	return string(m)
}
