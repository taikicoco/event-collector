package model

import (
	"time"
)

type LogDetail struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	LogNameID uint      `json:"log_name_id"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
