package storage

import (
	"context"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
)

func (s *Storage) AddTeam(req aplication.AddTeamRequest) error {
	tx, err := s.pool.Begin(context.Background())
	if err != nil {
		s.WriteError(err)
		return ErrFailedToCreateTransaction
	}
	defer func() {
		if err := tx.Rollback(context.Background()); err != nil {
			s.WriteError(err)
		}
	}()
	for _, user := range req.Users {
		_, err = tx.Exec(context.Background(),
			insertUser, user.Id, user.Name, user.IsActive,
		)
		if err != nil {
			return err
		}
	}
	for _, user := range req.Users {
		_, err = tx.Exec(context.Background(),
			insertTeam, req.TeamName, user.Id)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		s.WriteError(err)
		return ErrCommitTransaction
	}
	return nil
}
