package application

import (
	"context"
	"server/api/domain/model"
	"server/api/domain/repository"
)

type GetteWebLogRequest struct {
	ID uint `json:"id"`
}

type CreateWebLogRequest struct {
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}

type UpdateWebLogRequest struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}

type WebLogApplicationInterface interface {
	GetWebLogs(context.Context) ([]*model.WebLog, error)
	GetWebLog(context.Context, uint) (*model.WebLog, error)
	CreateWebLog(context.Context, *CreateWebLogRequest) (*model.WebLog, error)
	UpdateWebLog(context.Context, *UpdateWebLogRequest) (*model.WebLog, error)
}

type webLogApplication struct {
	webLogRepo repository.WebLogRepository
}

func NewWebLogApplication(webLogRepo repository.WebLogRepository) WebLogApplicationInterface {
	return &webLogApplication{webLogRepo: webLogRepo}
}

func (wla *webLogApplication) GetWebLogs(ctx context.Context) ([]*model.WebLog, error) {
	return wla.webLogRepo.FindAll()
}

func (wla *webLogApplication) GetWebLog(ctx context.Context, id uint) (*model.WebLog, error) {
	return wla.webLogRepo.FindByID(id)
}

func (wla *webLogApplication) CreateWebLog(ctx context.Context, req *CreateWebLogRequest) (*model.WebLog, error) {
	webLog := &model.WebLog{
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	}

	createWebLog, err := wla.webLogRepo.Create(webLog)
	if err != nil {
		return nil, err
	}
	return createWebLog, nil
}

func (wla *webLogApplication) UpdateWebLog(ctx context.Context, req *UpdateWebLogRequest) (*model.WebLog, error) {
	webLog := &model.WebLog{
		ID:      req.ID,
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	}

	updateWebLog, err := wla.webLogRepo.Update(webLog)
	if err != nil {
		return nil, err
	}
	return updateWebLog, nil
}
