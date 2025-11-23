package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/request"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (s *Storage) DeactivateUsers(req aplication.DeactivateUsersRequest) ([]request.Request, error) {
	activeUsers, err := s.selectUsersNotFromTeam(req.TeamName)
	if err != nil {
		return nil, err
	}

	rows, err := s.pool.Query(context.Background(), selectReviwerData, req.TeamName)
	if err != nil {
		s.WriteError(err)
		return nil, ErrSendQuery
	}
	defer rows.Close()

	reassignData := make([]struct {
		id       vo.PullRequestId
		reviewer vo.UserId
	}, 0)

	var reviewer string
	var id vo.PullRequestId

	for rows.Next() {
		if err := rows.Scan(&id, &reviewer); err != nil {
			s.WriteError(err)
			return nil, ErrConvertResponce
		}

		reassignData = append(reassignData, struct {
			id       vo.PullRequestId
			reviewer vo.UserId
		}{
			id:       vo.PullRequestId(id),
			reviewer: vo.UserId(reviewer),
		})
	}

	prs := make(map[vo.PullRequestId][]vo.UserId)
	for _, val := range reassignData {
		prs[val.id] = append(prs[val.id], val.reviewer)
	}

	resp := make([]request.Request, 0)
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		s.WriteError(err)
		return nil, ErrCreateTransaction
	}
	defer func() {
		if err := tx.Rollback(context.Background()); err != nil {
			s.WriteError(err)
		}
	}()

	for prID, reviewers := range prs {
		if err := s.updateReviewers(tx, prID, reviewers, activeUsers); err != nil {
			return nil, err
		}

		resp = append(resp, request.Request{
			PullRequestId: prID,
			Reviewer:      reviewers,
		})
	}

	if err := tx.Commit(context.Background()); err != nil {
		s.WriteError(err)
		return nil, ErrCommitTransaction
	}
	return resp, nil
}
