package model

import (
	"time"
)

type WebLogData struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	WebLogID   uint      `json:"web_log_id"`
	Access     uint      `json:"access"`
	Conversion uint      `json:"conversion"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
