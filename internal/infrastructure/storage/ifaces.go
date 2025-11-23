package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Logger interface {
	WriteError(error)
	WriteStatus(string)
	WriteDebugStatus(string)
}

type PGXPoolInterface interface { // Интерфейс будет полезен для юнит тестов (но на них определенно не хватит времени)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Ping(context.Context) error
	Close()
}
