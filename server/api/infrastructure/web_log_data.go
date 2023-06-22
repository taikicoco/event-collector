package repository

import (
	"context"
	"server/api/domain/model"
)

type WebLogDataRepositoryInterface interface {
	Update(context.Context, *model.WebLogData) (*model.WebLogData, error)
}

type webLogDataRepository struct{}

func NewWebLogDataRepository() WebLogDataRepositoryInterface {
	return &webLogDataRepository{}
}

func (wldr *webLogDataRepository) Update(ctx context.Context, webLogData *model.WebLogData) (*model.WebLogData, error) {
	db := dbInit()
	wld := db.Create(webLogData)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLogData, nil
}
