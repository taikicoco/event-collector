package model

import (
	"time"
)

type LogDetail struct {
	ID        uint      `json:"id" db:"id"`
	LogNameID uint      `json:"logNameId" db:"log_name_id"`
	LogName   *LogName  `json:"logName" db:"log_names"`
	Version   uint      `json:"version" db:"version"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
