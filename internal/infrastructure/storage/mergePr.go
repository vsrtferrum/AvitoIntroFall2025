package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
)

func (s *Storage) MergePr(req aplication.MergePullRequestRequest) error {
	_, err := s.pool.Exec(context.Background(), mergePr, req.PullRequestId)
	if err != nil {
		s.WriteError(err)
		return ErrSendQuery
	}
	return nil
}
