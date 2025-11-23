package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (s *Storage) GetReviwerStats(req aplication.StatReviewRequest) ([]aplication.StatReviewResponce, error) {
	rows, err := s.pool.Query(context.Background(), selectStat, req.Cursor, req.Limit)
	if err != nil {
		s.WriteError(err)
		return nil, ErrSendQuery
	}
	defer rows.Close()
	resp := make([]aplication.StatReviewResponce, 0)
	var id vo.UserId
	var reviewedPr int
	for rows.Next() {
		if err := rows.Scan(&id, &reviewedPr); err != nil {
			s.WriteError(err)
			return nil, ErrConvertResponce
		}

		resp = append(resp, aplication.StatReviewResponce{
			UserId:     id,
			ReviewedPr: reviewedPr,
		})
	}
	return resp, nil
}
