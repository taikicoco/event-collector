package application

import (
	"context"
	"server/api/domain/model"
	"server/api/repository"
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

type webLogApplication struct{}

func NewWebLogApplication() WebLogApplicationInterface {
	return &webLogApplication{}
}

func (wla *webLogApplication) GetWebLogs(ctx context.Context) ([]*model.WebLog, error) {
	wlr := repository.NewWebLogRepository()
	return wlr.FindAll(ctx)
}

func (wla *webLogApplication) GetWebLog(ctx context.Context, id uint) (*model.WebLog, error) {
	wlr := repository.NewWebLogRepository()
	return wlr.FindByID(ctx, id)
}

func (wla *webLogApplication) CreateWebLog(ctx context.Context, req *CreateWebLogRequest) (*model.WebLog, error) {
	webLog := &model.WebLog{
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	}

	wlr := repository.NewWebLogRepository()
	return wlr.Create(ctx, webLog)
}

func (wla *webLogApplication) UpdateWebLog(ctx context.Context, req *UpdateWebLogRequest) (*model.WebLog, error) {
	webLog := &model.WebLog{
		ID:      req.ID,
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	}
	wlr := repository.NewWebLogRepository()
	return wlr.Update(ctx, webLog)
}
