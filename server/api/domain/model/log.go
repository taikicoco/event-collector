package model

import (
	"time"
)

type Log struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	LogDetailID uint      `json:"log_detail_id"`
	Access      uint      `json:"access"`
	Conversion  uint      `json:"conversion"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
