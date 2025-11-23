package aplication

import (
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/team"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (g *GitModel) AddTeam(req AddTeamRequest) (*AddTeamResponce, error) {
	err := g.storage.AddTeam(req)
	if err != nil {
		g.logger.WriteError(err)
		return nil, ErrFailedToAddTeam
	}
	g.usersMtx.Lock()
	g.teamsMtx.Lock()
	defer g.usersMtx.Unlock()
	defer g.teamsMtx.Unlock()
	userIds := make([]vo.UserId, 0, len(req.Users))
	for _, u := range req.Users {
		userIds = append(userIds, u.Id)
	}

	g.teams[req.TeamName] = *team.NewTeam(req.TeamName, userIds...)
	return &AddTeamResponce{
		TeamName: req.TeamName,
		Users:    req.Users,
	}, nil
}
