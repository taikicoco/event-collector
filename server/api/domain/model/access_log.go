package model

import (
	"time"
)

type AccessLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AccessLogDetailID uint `json:"access_log_detail_id"`
	Access   uint      `json:"access"`
	Conversion uint     `json:"conversion"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}