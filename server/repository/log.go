package repository

import (
	"context"

	"server/model"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type LogRepository struct{}

func (lr *LogRepository) Create(ctx context.Context, db *sqlx.DB, logDetailID uint) (*model.Log, error) {
	log := &model.Log{}

	res, err := db.ExecContext(ctx,
		`
		insert into logs (log_detail_id, created_at) 
		values (?, CURRENT_TIMESTAMP)
		`,
		logDetailID)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.ID = uint(lastInsertId)
	log.LogDetailID = logDetailID

	return log, nil
}
