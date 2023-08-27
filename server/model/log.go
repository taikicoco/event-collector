package model

import (
	"time"
)

type Log struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	LogDetailID uint      `json:"log_detail_id"`
	CreatedAt   time.Time `json:"createdAt"`
}
