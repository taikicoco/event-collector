package application

import (
	"context"
	"server/api/domain/model"
)

type CreateLogRequest struct {
	LogDetailID uint `json:"id" param:"id"`
	Access      uint `json:"access"`
	Conversion  uint `json:"conversion"`
}

func CreateLog(ctx context.Context, req *CreateLogRequest) (*model.Log, error) {

	log := &model.Log{
		LogDetailID: req.LogDetailID,
		Access:      req.Access,
		Conversion:  req.Conversion,
	}

	db := dbInit()
	l := db.Create(log)
	err := l.Error
	if err != nil {
		return nil, err
	}
	return log, nil
}
