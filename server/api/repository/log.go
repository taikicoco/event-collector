package repository

import (
	"context"
	"server/api/domain/model"
)

type LogRepository interface {
	Create(context.Context, *model.Log) (*model.Log, error)
}

type logRepository struct{}

func NewLogRepository() LogRepository {
	return &logRepository{}
}

func(lr *logRepository) Create(ctx context.Context, log *model.Log) (*model.Log, error) {
	db := dbInit()
	l := db.Create(log)
	err := l.Error
	if err != nil {
		return nil, err
	}
	return log, nil
}
	
