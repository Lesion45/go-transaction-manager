package service

import (
	"context"
	"go-transaction-manager/internal/entity"
	"go-transaction-manager/internal/repository"
)

type ReservationService struct {
	reservationRepo repository.Reservation
}

func NewReservationService(reservationRepo repository.Reservation) *ReservationService {
	return &ReservationService{reservationRepo: reservationRepo}
}

func (s *ReservationService) ReserveBalance(ctx context.Context, input ReservationReserveBalanceInput) error {
	reservation := entity.Reservation{
		OrderID:   input.OrderID,
		UserID:    input.UserID,
		ServiceID: input.ServiceID,
		Amount:    input.Amount,
	}

	return s.reservationRepo.ReserveBalance(ctx, reservation)
}

func (s *ReservationService) CommitReservedBalance(ctx context.Context, input ReservationCommitReservedBalanceInput) error {
	reservation := entity.Reservation{
		OrderID:   input.OrderID,
		UserID:    input.UserID,
		ServiceID: input.ServiceID,
		Amount:    input.Amount,
	}

	return s.reservationRepo.CommitReservedBalance(ctx, reservation)
}
