package model

import (
	"time"
)

type WebLog struct {
	ID         uint          `json:"id" gorm:"primaryKey"`
	Name       string        `json:"name"`
	Version    uint          `json:"version"`
	PageUrl    string        `json:"page_url"`
	WebLogData []*WebLogData `json:"web_log_data"`
	CreatedAt  time.Time     `json:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt"`
}
