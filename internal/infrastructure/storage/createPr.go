package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
)

func (s *Storage) CreatePr(req aplication.CreatePullRequestRequest) (*aplication.CreatePullRequestResponce, error) {
	activeUsers, err := s.SelectActiveUsers()
	if err != nil {
		return nil, err
	}
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		s.WriteError(err)
		return nil, ErrFailedToCreateTransaction
	}
	defer func() {
		if err := tx.Rollback(context.Background()); err != nil {
			s.WriteError(err)
		}
	}()

	if _, err = tx.Exec(context.Background(), createPr, req.PullRequestId, req.PullRequestName, req.AuthorId, true); err != nil {
		s.WriteError(err)
		return nil, ErrExecTransaction
	}
	for _, user := range activeUsers {
		if _, err := tx.Exec(context.Background(), addReviewers, req.PullRequestId, user); err != nil {
			s.WriteError(err)
			return nil, ErrExecTransaction
		}
	}
	if err := tx.Commit(context.Background()); err != nil {
		s.WriteError(err)
		return nil, ErrCommitTransaction
	}
	return &aplication.CreatePullRequestResponce{
		AssignedReviewers: activeUsers,
	}, nil
}
