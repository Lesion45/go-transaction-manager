package service

import (
	"context"
	"github.com/google/uuid"
	"go-transaction-manager/internal/entity"
	"go-transaction-manager/internal/repository"
)

type UserAddBalanceInput struct {
	UserID uuid.UUID
	Amount float64
}

type UserGetBalanceInput struct {
	UserID uuid.UUID
}

type User interface {
	AddBalance(ctx context.Context, user UserAddBalanceInput) error
	GetBalance(ctx context.Context, user UserGetBalanceInput) (entity.Balance, error)
}

type ReservationReserveBalanceInput struct {
	OrderID   uuid.UUID
	UserID    uuid.UUID
	ServiceID uuid.UUID
	Amount    float64
	Info      string
}

type ReservationCommitReservedBalanceInput struct {
	OrderID   uuid.UUID
	UserID    uuid.UUID
	ServiceID uuid.UUID
	Amount    float64
	Info      string
}

type Reservation interface {
	ReserveBalance(ctx context.Context, reservation ReservationReserveBalanceInput) error
	CommitReservedBalance(ctx context.Context, reservation ReservationCommitReservedBalanceInput) error
}

type Services struct {
	User        User
	Reservation Reservation
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(dependencies ServicesDependencies) *Services {
	return &Services{
		User:        NewUserService(dependencies.Repos.User),
		Reservation: NewReservationService(dependencies.Repos.Reservation),
	}
}
