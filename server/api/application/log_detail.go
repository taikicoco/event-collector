package application

import (
	"context"
	"server/api/domain/model"
	"server/api/repository"
)

type CreateLogDetailRequest struct {
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}

type GetLogDetailRequest struct {
	ID uint `json:"id" param:"id"`
}

func CreateLogDetail(ctx context.Context, req *CreateLogDetailRequest) (*model.LogDetail, error) {
	logDetail := &model.LogDetail{
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	}
	ldr := repository.NewLogDetailRepository()
	return ldr.Create(ctx, logDetail)
}

func GetLogDetail(ctx context.Context, id uint) (*model.LogDetail, error) {
	ldr := repository.NewLogDetailRepository()
	return ldr.Get(ctx, id)
}
