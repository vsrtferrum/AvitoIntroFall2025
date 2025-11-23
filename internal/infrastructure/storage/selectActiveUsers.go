package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (s *Storage) SelectActiveUsers() ([]vo.UserId, error) {
	rows, err := s.pool.Query(context.Background(), selectTwoFreeActiveUsers)
	if err != nil {
		s.WriteError(err)
		return nil, ErrSendQuery
	}
	defer rows.Close()
	activeUsers := make([]vo.UserId, 0, 2)
	var id string
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			s.WriteError(err)
			return nil, ErrConvertResponce
		}

		activeUsers = append(activeUsers, vo.UserId(id))
	}
	return activeUsers, nil
}
