package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
)

func (s *Storage) SetUserActive(req aplication.SetIsAtivateRequest) error {
	_, err := s.pool.Exec(context.Background(), setUserActive, req.IsActive, req.UserId)
	if err != nil {
		s.WriteError(err)
		return ErrSendQuery
	}
	return nil
}
