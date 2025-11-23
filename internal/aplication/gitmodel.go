package aplication

import (
	"sync"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/request"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/team"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/user"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

type GitModel struct {
	logger      Logger
	storage     Storage
	users       map[vo.UserId]user.User
	teams       map[vo.TeamName]team.Team
	requests    map[vo.PullRequestId]request.Request
	usersMtx    sync.Mutex
	teamsMtx    sync.Mutex
	requestsMtx sync.Mutex
}

func NewGitModel(users []user.User, teams []team.Team, requests []request.Request, storage Storage, loggger Logger) *GitModel {
	usersMap := make(map[vo.UserId]user.User)
	for _, u := range users {
		usersMap[u.Id] = u
	}

	teamMap := make(map[vo.TeamName]team.Team)
	for _, t := range teams {
		teamMap[t.TeamName] = t
	}

	requestsMap := make(map[vo.PullRequestId]request.Request)
	for _, r := range requests {
		requestsMap[r.PullRequestId] = r
	}
	return &GitModel{
		usersMtx:    sync.Mutex{},
		teamsMtx:    sync.Mutex{},
		requestsMtx: sync.Mutex{},
		logger:      loggger,
		users:       usersMap,
		teams:       teamMap,
		requests:    requestsMap,
		storage:     storage,
	}
}
