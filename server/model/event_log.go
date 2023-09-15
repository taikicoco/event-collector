package model

import (
	"time"
)

type EventLog struct {
	ID        uint      `json:"id" db:"id"`
	EventID   uint      `json:"event_id" db:"event_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
