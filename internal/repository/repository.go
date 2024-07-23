package repository

import (
	"context"
	"github.com/google/uuid"
	"go-transaction-manager/internal/entity"
	"go-transaction-manager/internal/repository/pgdb"
	"go-transaction-manager/pkg/postgres"
)

type User interface {
	AddBalance(ctx context.Context, id uuid.UUID, amount float64) error
	GetBalance(ctx context.Context, id uuid.UUID) (entity.Balance, error)
}

type Reservation interface {
	ReserveBalance(ctx context.Context, user entity.Reservation) error
	CommitReservedBalance(ctx context.Context, user entity.Reservation) error
}

type Repositories struct {
	User
	Reservation
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User:        pgdb.NewUserRepository(pg),
		Reservation: pgdb.NewReservationRepository(pg),
	}
}
