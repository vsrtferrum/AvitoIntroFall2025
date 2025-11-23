package aplication

import (
	"slices"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (g *GitModel) SetActivate(userId vo.UserId, flag bool) (*SetIsAtivateResponce, error) {
	err := g.storage.SetUserActive(SetIsAtivateRequest{
		UserId:   userId,
		IsActive: flag,
	})
	if err != nil {
		g.logger.WriteError(err)
		return nil, ErrUserNotFound
	}

	g.usersMtx.Lock()
	defer g.usersMtx.Unlock()
	usr, ok := g.users[userId]
	if !ok {
		return nil, ErrUserNotFound
	}
	usr.IsActive = flag
	g.users[userId] = usr

	return &SetIsAtivateResponce{
		UserId:   userId,
		UserName: usr.Name,
		IsActive: flag,
		TeamName: func() vo.TeamName {
			for team, val := range g.teams {
				if slices.Contains(val.Get(), userId) {
						return team
					}
			}
			return ""
		}(),
	}, nil
}
