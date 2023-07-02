package infrastructure

import (
	"gorm.io/gorm"
	"server/api/domain/model"
	"server/api/domain/repository"
)

type webLogDataRepository struct {
	Conn *gorm.DB
}

func NewWebLogDataRepository(conn *gorm.DB) repository.WebLogDataRepository {
	return &webLogDataRepository{Conn: conn}
}

func (wldr *webLogDataRepository) Update(webLogData *model.WebLogData) (*model.WebLogData, error) {
	db := repository.DBInit()
	wld := db.Create(webLogData)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLogData, nil
}
