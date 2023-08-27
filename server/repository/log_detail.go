package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type LogDetailRepository struct{}

func (lr *LogDetailRepository) GetByID(ctx context.Context, db *sqlx.DB, id uint) (*model.LogDetail, error) {
	logDetail := &model.LogDetail{}
	err := db.Get(logDetail, `select * from log_details where id = ?`, id)

	logDetail.LogName, err = (&LogNameRepository{}).GetByID(ctx, db, logDetail.LogNameID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return logDetail, nil
}
