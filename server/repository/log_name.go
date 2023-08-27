package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type LogNameRepository struct{}

func (lr *LogNameRepository) GetByID(ctx context.Context, db *sqlx.DB, id uint) (*model.LogName, error) {
	logName := &model.LogName{}
	err := db.Get(logName, `select * from log_names where id = ?`, id)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	return logName, nil
}
