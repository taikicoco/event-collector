package model

type LogName struct {
	ID   uint   `json:"id" db:"primaryKey"`
	Name string `json:"name"`
}
