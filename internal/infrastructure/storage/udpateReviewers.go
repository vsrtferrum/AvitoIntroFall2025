package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (s *Storage) updateReviewers(tx pgx.Tx, prID vo.PullRequestId, oldReviewers []vo.UserId, newReviewers []vo.UserId) error {
	for i, oldReviewer := range oldReviewers {
		if _, err := tx.Exec(context.Background(), addReviewers, prID, newReviewers[i]); err != nil {
			s.WriteError(err)
			return ErrExecTransaction
		}
		if _, err := tx.Exec(context.Background(), dropUserFromReview, prID, oldReviewer); err != nil {
			s.WriteError(err)
			return ErrExecTransaction
		}
	}
	return nil
}
