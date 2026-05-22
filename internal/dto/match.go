package dto

import (
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/model"
)

func (r CreateMatchRequest) ToModel(matchDate time.Time) *model.Match {
	return &model.Match{MatchDate: matchDate, MatchTime: r.MatchTime, HomeTeamID: r.HomeTeamID, AwayTeamID: r.AwayTeamID, Status: model.MatchStatusScheduled}
}

type CreateMatchRequest struct {
	MatchDate  string `json:"match_date" binding:"required"`
	MatchTime  string `json:"match_time" binding:"required"`
	HomeTeamID int64  `json:"home_team_id" binding:"required"`
	AwayTeamID int64  `json:"away_team_id" binding:"required"`
}

type UpdateMatchRequest struct {
	MatchDate  string `json:"match_date" binding:"omitempty"`
	MatchTime  string `json:"match_time" binding:"omitempty"`
	HomeTeamID int64  `json:"home_team_id" binding:"omitempty"`
	AwayTeamID int64  `json:"away_team_id" binding:"omitempty"`
}

type GoalInput struct {
	PlayerID     int64 `json:"player_id" binding:"required"`
	MinuteScored int64 `json:"minute" binding:"required,min=1,max=130"`
}

type MatchResultRequest struct {
	HomeScore int64       `json:"home_score" binding:"min=0"`
	AwayScore int64       `json:"away_score" binding:"min=0"`
	Goals     []GoalInput `json:"goals"`
}

type GoalResponse struct {
	ID           int64  `json:"id"`
	PlayerID     int64  `json:"player_id"`
	PlayerName   string `json:"player_name"`
	MinuteScored int64  `json:"minute"`
}

type MatchResponse struct {
	ID        int64          `json:"id"`
	MatchDate string         `json:"match_date"`
	MatchTime string         `json:"match_time"`
	HomeTeam  *TeamResponse  `json:"home_team"`
	AwayTeam  *TeamResponse  `json:"away_team"`
	HomeScore int64          `json:"home_score"`
	AwayScore int64          `json:"away_score"`
	Status    string         `json:"status"`
	Goals     []GoalResponse `json:"goals,omitempty"`
	CreatedAt string         `json:"created_at"`
}


type TopScorer struct {
	PlayerID int64  `json:"player_id"`
	Name     string `json:"name"`
	Goals    int64  `json:"goals"`
}

type ReportResponse struct {
	Match                           MatchResponse `json:"match"`
	FinalScore                      FinalScore    `json:"final_score"`
	ResultStatus                    string        `json:"result_status"`
	TopScorer                       *TopScorer    `json:"top_scorer"`
	HomeTeamTotalWinsUntilThisMatch int64         `json:"home_team_total_wins_until_this_match"`
	AwayTeamTotalWinsUntilThisMatch int64         `json:"away_team_total_wins_until_this_match"`
}

type FinalScore struct {
	Home int64 `json:"home"`
	Away int64 `json:"away"`
}
