package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Log struct {
	db      *sqlx.DB
	LogRepo *repository.LogRepository
}

func NewLog(db *sqlx.DB) *Log {
	return &Log{
		db: db,
	}
}

func (l *Log) Create(ctx context.Context, id uint) (*model.Log, error) {
	return l.LogRepo.Create(ctx, l.db, id)
}
