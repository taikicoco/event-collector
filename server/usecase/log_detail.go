package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type LogDetail struct {
	db          *sqlx.DB
	logDetailRepo *repository.LogDetailRepository
}

func NewLogDetail(db *sqlx.DB) *LogDetail {
	return &LogDetail{
		db: db,
	}
}

func (l *LogDetail) GetByID(ctx context.Context, id uint) (*model.LogDetail, error) {
	return l.logDetailRepo.GetByID(ctx, l.db, id)
}
