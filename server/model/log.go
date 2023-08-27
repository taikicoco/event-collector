package model

import (
	"time"
)

type Log struct {
	ID          uint      `json:"id" db:"id"`
	LogDetailID uint      `json:"log_detail_id" db:"log_detail_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}
