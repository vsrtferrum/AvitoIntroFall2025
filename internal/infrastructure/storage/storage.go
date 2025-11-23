package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Logger
	pool PGXPoolInterface
}

func NewStorage(ctx context.Context, cfg *Config, logger Logger) (*Storage, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s ",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	poolCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.WriteError(err)
		return nil, ErrCreateConnection
	}

	if cfg.PoolSize > 0 {
		poolCfg.MaxConns = cfg.PoolSize
	}

	poolCfg.HealthCheckPeriod = cfg.HealthCheckPeriod
	poolCfg.MaxConnLifetime = cfg.MaxConnLifetime
	poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		logger.WriteError(err)
		return nil, ErrCreateConnection
	}
	if err = pool.Ping(ctx); err != nil {
		logger.WriteError(err)
		return nil, ErrCreateConnection
	}

	return &Storage{pool: pool, Logger: logger}, nil
}
