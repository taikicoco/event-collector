package application

import (
	"context"
	"server/api/domain/model"
	"server/api/domain/repository"
)

type Application struct {
	*ApplicationBundle
}

type ApplicationBundle struct {
	webLogRepository     repository.WebLogRepository
	webLogDataRepository repository.WebLogDataRepository
}

type ApplicationInterface interface {
	// WebLog
	GetWebLogs(context.Context) ([]*model.WebLog, error)
	GetWebLog(context.Context, uint) (*model.WebLog, error)
	CreateWebLog(context.Context, *CreateWebLogRequest) (*model.WebLog, error)
	UpdateWebLog(context.Context, *UpdateWebLogRequest) (*model.WebLog, error)

	// WebLogData
	UpdateWebLogData(context.Context, *UpdateWebLogDataRequest) (*model.WebLogData, error)
}
