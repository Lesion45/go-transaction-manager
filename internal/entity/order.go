package entity

import "github.com/google/uuid"

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ServiceID uuid.UUID
	Amount    float64
	Info      string
}
