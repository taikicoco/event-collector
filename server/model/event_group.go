package model

type EventGroup struct {
	ID        uint   `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	DomainUrl string `json:"domainUrl" db:"domain_url"`
}
