package model

import (
	"time"
)

type Log struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Log       uint      `json:"log"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
