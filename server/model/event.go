package model

type Event struct {
	ID           uint   `json:"id" db:"id"`
	EventGroupID uint   `json:"event_group_id" db:"event_group_id"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
}
