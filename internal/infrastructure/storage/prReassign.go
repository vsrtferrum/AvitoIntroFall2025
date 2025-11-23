package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (s *Storage) PrReassign(req aplication.ReassignPullRequestRequest) (*aplication.ReassignPullRequestResponce, error) {
	activeUsers, err := s.SelectActiveUsers()
	if err != nil {
		return nil, err
	}

	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		s.WriteError(err)
		return nil, ErrCreateTransaction
	}
	err = s.updateReviewers(tx, req.PullRequestId, []vo.UserId{req.OldReviewerId}, activeUsers[:1])
	if err != nil {
		s.WriteError(err)
		return nil, ErrCommitTransaction
	}
	return &aplication.ReassignPullRequestResponce{
		AssignedReviewers: activeUsers,
	}, nil

}
