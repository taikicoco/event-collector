package model

import (
	"time"
)

type AccessLogDetail struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Version   uint      `json:"version"`
	PageUrl   string    `json:"page_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
