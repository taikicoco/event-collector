package model

import (
	"time"

	"gorm.io/gorm"
)

type Layout struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	LogKey    string         `json:"log_key"`
	LogValue  string         `json:"log_value"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
