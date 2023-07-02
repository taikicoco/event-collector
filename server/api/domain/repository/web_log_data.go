package repository

import (
	"server/api/domain/model"
)

type WebLogDataRepository interface {
	Update(*model.WebLogData) (*model.WebLogData, error)
}
