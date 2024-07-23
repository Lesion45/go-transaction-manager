package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Balance  float64
	Reserved float64
}
