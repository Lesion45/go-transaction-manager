package models

import "github.com/google/uuid"

type UserAddBalance struct {
	UserID uuid.UUID `json:"user-id"`
	Amount float64   `json:"amount"`
}

type ReservationReserveBalance struct {
	OrderID   uuid.UUID `json:"order_id"`
	UserID    uuid.UUID `json:"user-id"`
	ServiceID uuid.UUID `json:"service-id"`
	Amount    float64   `json:"amount"`
	Info      string    `json:"info"`
}

type ReservationCommitReservedBalance struct {
	OrderID   uuid.UUID `json:"order_id"`
	UserID    uuid.UUID `json:"user-id"`
	ServiceID uuid.UUID `json:"service-id"`
	Amount    float64   `json:"amount"`
	Info      string    `json:"info"`
}

type UserGetBalance struct {
	UserID uuid.UUID `json:"user-id"`
}
