package application

import (
	"net/http"
	"server/api/domain/model"

	"github.com/labstack/echo/v4"
)

type createLogRequest struct {
	LogDetailID uint `json:"id" param:"id"`
	Access      uint `json:"access"`
	Conversion  uint `json:"conversion"`
}

func CreateLog(e echo.Context) error {
	var rl createLogRequest
	if err := e.Bind(&rl); err != nil {
		return err
	}

	log := model.Log{
		LogDetailID: rl.LogDetailID,
		Access:      rl.Access,
		Conversion:  rl.Conversion,
	}

	db := dbInit()
	l := db.Create(&log)
	err := l.Error
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, log)
}
