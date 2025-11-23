package aplication

import (
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/team"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (g *GitModel) GetTeam(team vo.TeamName) (*team.Team, error) {
	g.teamsMtx.Lock()
	defer g.teamsMtx.Unlock()
	t, ok := g.teams[team]
	if ok {
		return &t, nil
	}
	return nil, ErrNoTeam
}
