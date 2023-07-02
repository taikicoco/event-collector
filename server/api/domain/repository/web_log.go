package repository

import (
	"server/api/domain/model"
)

type WebLogRepository interface {
	Create(*model.WebLog) (*model.WebLog, error)
	Update(*model.WebLog) (*model.WebLog, error)
	FindAll() ([]*model.WebLog, error)
	FindByID(uint) (*model.WebLog, error)
}
