package application

import (
	"net/http"
	"server/api/domain/model"

	"github.com/labstack/echo/v4"
)

type createLogDetailRequest struct {
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}

type getLogDetailRequest struct {
	ID uint `json:"id" param:"id"`
}

func CreateLogDetail(e echo.Context) error {
	var r createLogDetailRequest
	if err := e.Bind(&r); err != nil {
		return err
	}

	logDetail := model.LogDetail{
		Name:    r.Name,
		Version: r.Version,
		PageUrl: r.PageUrl,
	}

	db := dbInit()
	l := db.Create(&logDetail)
	err := l.Error
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, logDetail)
}

func GetLogDetail(e echo.Context) error {

	var r getLogDetailRequest
	if err := e.Bind(&r); err != nil {
		return err
	}

	db := dbInit()

	var logDetail model.LogDetail
	l := db.First(&logDetail, r.ID)
	err := l.Error
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, logDetail)
}