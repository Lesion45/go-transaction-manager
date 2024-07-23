package entity

import (
	"github.com/google/uuid"
	"time"
)

type Report struct {
	Date      time.Time
	ServiceID uuid.UUID
	UserID    uuid.UUID
	Amount    float64
	Info      string
}
