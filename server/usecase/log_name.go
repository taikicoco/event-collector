package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type LogName struct {
	db          *sqlx.DB
	logNameRepo *repository.LogNameRepository
}

func NewLogName(db *sqlx.DB) *LogName {
	return &LogName{
		db: db,
	}
}

func (l *LogName) Get(ctx context.Context, id uint) (*model.LogName, error) {
	return l.logNameRepo.GetByID(ctx, l.db, id)
}
