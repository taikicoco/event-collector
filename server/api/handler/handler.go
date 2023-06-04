package handler

import (
	"net/http"
	"server/api/domain/model"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AssignRoutes(e *echo.Echo) error {
	e.POST("/log/detail", createLogDetail)
	e.GET("/log/detail/:id", getLogDetail)
	e.POST("/log/:id", createLog)
	return nil
}

type createLogDetailRequest struct {
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}

type getLogDetailRequest struct {
	ID uint `json:"id" param:"id"`
}

type createLogRequest struct {
	LogDetailID uint `json:"id" param:"id"`
	Access      uint `json:"access"`
	Conversion  uint `json:"conversion"`
}

func createLogDetail(e echo.Context) error {
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

func getLogDetail(e echo.Context) error {

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

func createLog(e echo.Context) error {
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

func dbInit() *gorm.DB {
	dsn := "user:password@tcp(mysql)/db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
