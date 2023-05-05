package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    string         `json:"userId"`
	Pawssword string         `json:"password"`
	Comment   string         `json:"comment"`
	Nickname  string         `json:"nickname"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
