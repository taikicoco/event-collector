package model

type LogName struct {
	ID   uint   `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
