package model

type EventSource struct {
	ID        uint   `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	SourseUrl string `json:"sourseUrl" db:"sourse_url"`
}
