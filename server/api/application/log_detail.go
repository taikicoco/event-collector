package application

import (
	"context"
	"server/api/domain/model"
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

	db := dbInit()
	l := db.Create(logDetail)
	err := l.Error
	if err != nil {
		return nil, err
	}
	return logDetail, nil
}

func GetLogDetail(ctx context.Context, id uint) (*model.LogDetail, error) {
	db := dbInit()

	var logDetail *model.LogDetail
	l := db.First(&logDetail, id)
	err := l.Error
	if err != nil {
		return nil, err
	}
	return logDetail, nil
}
