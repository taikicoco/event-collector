package application

import (
	"server/api/infrastracture/persistence"
	"server/config"
)

type application struct {
	*ApplicationBundle
}

type ApplicationBundle struct {
	Repository   persistence.RepositoryInterface
	ServerConfig *config.ServerConfig
}

type ApplicationInterface interface {
}

func NewApplication(bdl *ApplicationBundle) ApplicationInterface {
	return &application{bdl}
}
