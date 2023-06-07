package repository

import (
	"context"
	"server/api/domain/model"
)

type LogDetailRepository interface {
	Create(context.Context, *model.LogDetail) (*model.LogDetail, error)
	Get(context.Context, uint) (*model.LogDetail, error)
}

type logDetailRepository struct{}

func NewLogDetailRepository() LogDetailRepository {
	return &logDetailRepository{}
}

func (ld *logDetailRepository) Create(ctx context.Context, logDetail *model.LogDetail) (*model.LogDetail, error) {
	db := dbInit()
	l := db.Create(logDetail)
	err := l.Error
	if err != nil {
		return nil, err
	}
	return logDetail, nil
}

func (ld *logDetailRepository) Get(ctx context.Context, id uint) (*model.LogDetail, error) {
	db := dbInit()

	var logDetail *model.LogDetail
	l := db.First(&logDetail, id)
	err := l.Error
	if err != nil {
		return nil, err
	}
	return logDetail, nil
}
