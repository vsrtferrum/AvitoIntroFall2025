package team

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

type Team struct {
	TeamName vo.TeamName
	users []vo.UserId
}

func NewTeam(teamName vo.TeamName, users ...vo.UserId) *Team {
	return &Team{
		TeamName: teamName,
		users:    users,
	}
}
