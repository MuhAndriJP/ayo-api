package dto

import (
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

func MatchListToResponse(matches []model.Match) []*MatchResponse {
	res := make([]*MatchResponse, len(matches))
	for i := range matches {
		res[i] = MatchToResponse(&matches[i])
	}
	return res
}

func ReportFromMatch(match *model.Match, homeWins, awayWins int64) *ReportResponse {
	r := &ReportResponse{
		Match:                           *MatchToResponse(match),
		ResultStatus:                    DetermineResultStatus(match.HomeScore, match.AwayScore),
		TopScorer:                       ComputeTopScorer(match.Goals),
		HomeTeamTotalWinsUntilThisMatch: homeWins,
		AwayTeamTotalWinsUntilThisMatch: awayWins,
	}
	r.FinalScore.Home = match.HomeScore
	r.FinalScore.Away = match.AwayScore
	return r
}

func MatchToResponse(m *model.Match) *MatchResponse {
	var goals []GoalResponse
	for _, g := range m.Goals {
		goals = append(goals, GoalResponse{
			ID:           int64(g.ID),
			PlayerID:     g.PlayerID,
			PlayerName:   g.Player.Name,
			MinuteScored: g.MinuteScored,
		})
	}
	return &MatchResponse{
		ID:        int64(m.ID),
		MatchDate: util.FormatDate(m.MatchDate),
		MatchTime: m.MatchTime,
		HomeTeam:  TeamToResponse(&m.HomeTeam),
		AwayTeam:  TeamToResponse(&m.AwayTeam),
		HomeScore: m.HomeScore,
		AwayScore: m.AwayScore,
		Status:    m.Status.String(),
		Goals:     goals,
		CreatedAt: util.FormatDateTime(m.CreatedAt),
	}
}

func DetermineResultStatus(home, away int64) string {
	switch {
	case home > away:
		return model.ResultHomeWin
	case away > home:
		return model.ResultAwayWin
	default:
		return model.ResultDraw
	}
}

func ComputeTopScorer(goals []model.Goal) *TopScorer {
	counts := map[int64]struct {
		name  string
		count int64
	}{}
	for _, g := range goals {
		entry := counts[g.PlayerID]
		entry.name = g.Player.Name
		entry.count++
		counts[g.PlayerID] = entry
	}
	var top *TopScorer
	for pid, v := range counts {
		if top == nil || v.count > top.Goals {
			top = &TopScorer{
				PlayerID: pid,
				Name:     v.name,
				Goals:    v.count,
			}
		}
	}
	return top
}
