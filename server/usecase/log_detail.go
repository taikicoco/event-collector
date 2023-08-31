package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type LogDetail struct {
	db            *sqlx.DB
	logDetailRepo *repository.LogDetailRepository
	logNameRepo   *repository.LogNameRepository
}

func NewLogDetail(db *sqlx.DB) *LogDetail {
	return &LogDetail{
		db: db,
	}
}

func (l *LogDetail) GetByID(ctx context.Context, id uint) (*model.LogDetail, error) {
	logDetail, err := l.logDetailRepo.GetByID(ctx, l.db, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	logName, err := l.logNameRepo.GetByID(ctx, l.db, logDetail.LogNameID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	logDetail.LogName = logName
	return logDetail, nil
}
