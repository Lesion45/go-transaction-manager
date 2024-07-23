package entity

import "github.com/google/uuid"

type Reservation struct {
	OrderID   uuid.UUID
	UserID    uuid.UUID
	ServiceID uuid.UUID
	Amount    float64
	Info      string
}
