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

func GetWebLogs(ctx context.Context) ([]*model.WebLog, error) {
	wlr := repository.NewWebLogRepository()
	return wlr.FindAll(ctx)
}

func GetWebLog(ctx context.Context, id uint) (*model.WebLog, error) {
	wlr := repository.NewWebLogRepository()
	return wlr.FindByID(ctx, id)
}

func CreateWebLog(ctx context.Context, req *CreateWebLogRequest) (*model.WebLog, error) {
	webLog := &model.WebLog{
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	}

	wlr := repository.NewWebLogRepository()
	return wlr.Create(ctx, webLog)
}

func UpdateWebLog(ctx context.Context, req *UpdateWebLogRequest) (*model.WebLog, error) {
	webLog := &model.WebLog{
		ID:      req.ID,
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	}
	wlr := repository.NewWebLogRepository()
	return wlr.Update(ctx, webLog)
}
