package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/request"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/team"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/user"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (s *Storage) GetAllUsers() ([]user.User, error) {
	rows, err := s.pool.Query(context.Background(), getAllUsersQuery)
	if err != nil {
		s.WriteError(err)
		return nil, ErrSendQuery
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var id, username string
		var isActive bool

		if err := rows.Scan(&id, &username, &isActive); err != nil {
			s.WriteError(err)
			return nil, ErrConvertResponce
		}

		users = append(users, user.User{
			Id:       vo.UserId(id),
			Name:     vo.UserName(username),
			IsActive: isActive,
		})
	}

	return users, nil
}

func (s *Storage) GetAllTeams() ([]team.Team, error) {
	rows, err := s.pool.Query(context.Background(), getAllTeamsQuery)
	if err != nil {
		s.WriteError(err)
		return nil, ErrSendQuery
	}
	defer rows.Close()

	teamMap := make(map[vo.TeamName][]vo.UserId)
	for rows.Next() {
		var teamName, userID string
		if err := rows.Scan(&teamName, &userID); err != nil {
			s.WriteError(err)
			return nil, ErrConvertResponce
		}

		teamNameVO := vo.TeamName(teamName)
		userIDVO := vo.UserId(userID)
		teamMap[teamNameVO] = append(teamMap[teamNameVO], userIDVO)
	}

	var teams []team.Team
	for teamName, userIDs := range teamMap {
		teams = append(teams, *team.NewTeam(teamName, userIDs...))
	}

	return teams, nil
}

func (s *Storage) GetAllRequests() ([]request.Request, error) {
	prRows, err := s.pool.Query(context.Background(), getAllPrsQuery)
	if err != nil {
		s.WriteError(err)
		return nil, ErrSendQuery
	}
	defer prRows.Close()

	prMap := make(map[vo.PullRequestId]request.Request)
	for prRows.Next() {
		var id, name, author string
		var opened bool

		if err := prRows.Scan(&id, &name, &author, &opened); err != nil {
			s.WriteError(err)
			return nil, ErrConvertResponce
		}

		status := vo.StatusOPEN
		if !opened {
			status = vo.StatusMERGED
		}

		prMap[vo.PullRequestId(id)] = request.Request{
			PullRequestId:   vo.PullRequestId(id),
			PullRequestName: vo.PullRequestName(name),
			Author:          vo.UserId(author),
			Status:          status,
			Reviewer:        []vo.UserId{},
		}
	}

	reviewRows, err := s.pool.Query(context.Background(), getAllReviewsQuery)
	if err != nil {
		s.WriteError(err)
		return nil, ErrSendQuery
	}
	defer reviewRows.Close()

	for reviewRows.Next() {
		var prID, reviewer string
		if err := reviewRows.Scan(&prID, &reviewer); err != nil {
			s.WriteError(err)
			return nil, ErrConvertResponce
		}

		prIDVO := vo.PullRequestId(prID)
		if pr, exists := prMap[prIDVO]; exists {
			pr.Reviewer = append(pr.Reviewer, vo.UserId(reviewer))
			prMap[prIDVO] = pr
		}
	}

	var requests []request.Request
	for _, req := range prMap {
		requests = append(requests, req)
	}

	return requests, nil
}

func (s *Storage) GetAllData() ([]user.User, []team.Team, []request.Request, error) {
	users, err := s.GetAllUsers()
	if err != nil {
		return nil, nil, nil, err
	}

	teams, err := s.GetAllTeams()
	if err != nil {
		return nil, nil, nil, err
	}

	requests, err := s.GetAllRequests()
	if err != nil {
		return nil, nil, nil, err
	}

	return users, teams, requests, nil
}
