package infrastructure

import (
	"gorm.io/gorm"
	"server/api/domain/model"
	"server/api/domain/repository"
)

type webLogRepository struct {
	Conn *gorm.DB
}

func NewWebLogRepository(conn *gorm.DB) repository.WebLogRepository {
	return &webLogRepository{Conn: conn}
}

func (wlr *webLogRepository) Create(webLog *model.WebLog) (*model.WebLog, error) {
	db := repository.DBInit()
	wld := db.Create(webLog)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLog, nil
}

func (wlr *webLogRepository) Update(webLog *model.WebLog) (*model.WebLog, error) {
	db := repository.DBInit()
	wld := db.Create(webLog)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLog, nil
}

func (wlr *webLogRepository) FindAll() ([]*model.WebLog, error) {
	db := repository.DBInit()
	var webLogs []*model.WebLog
	wld := db.Find(&webLogs)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return webLogs, nil
}

func (wlr *webLogRepository) FindByID(id uint) (*model.WebLog, error) {
	db := repository.DBInit()
	var webLog model.WebLog
	wld := db.First(&webLog, id)
	err := wld.Error
	if err != nil {
		return nil, err
	}
	return &webLog, nil
}
