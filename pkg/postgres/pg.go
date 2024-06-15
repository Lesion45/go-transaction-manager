package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	DB *pgxpool.Pool
}

var (
	pgInstance *Postgres
)

func NewPG(ctx context.Context, DSN string) (*Postgres, error) {
	const op = "storage.postgres.NewPG"
	db, err := pgxpool.New(ctx, DSN)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pgInstance = &Postgres{DB: db}

	return pgInstance, nil
}
