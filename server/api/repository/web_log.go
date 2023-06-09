package repository

import (
	"context"
	"server/api/domain/model"
)

type WebLogRepositoryInterface interface {
	Create(context.Context, *model.WebLog) (*model.WebLog, error)
	Update(context.Context, *model.WebLog) (*model.WebLog, error)
	FindAll(context.Context) ([]*model.WebLog, error)
	FindByID(context.Context, uint) (*model.WebLog, error)
}

type webLogRepository struct{}

func NewWebLogRepository() WebLogRepositoryInterface {
	return &webLogRepository{}
}

func (wlr *webLogRepository) Create(ctx context.Context, webLog *model.WebLog) (*model.WebLog, error) {
	db := dbInit()
	wld := db.Create(webLog)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLog, nil
}

func (wlr *webLogRepository) Update(ctx context.Context, webLog *model.WebLog) (*model.WebLog, error) {
	db := dbInit()
	wld := db.Create(webLog)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLog, nil
}

func (wlr *webLogRepository) FindAll(ctx context.Context) ([]*model.WebLog, error) {
	db := dbInit()
	var webLogs []*model.WebLog
	wld := db.Find(&webLogs)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLogs, nil
}

func (wlr *webLogRepository) FindByID(ctx context.Context, id uint) (*model.WebLog, error) {
	db := dbInit()
	var webLog model.WebLog
	wld := db.First(&webLog, id)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return &webLog, nil
}
