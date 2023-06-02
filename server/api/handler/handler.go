package handler

import (
	"net/http"
	"server/api/domain/model"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AssignRoutes(e *echo.Echo) error {
	e.GET("/log", getData)
	e.POST("/log", createData)
	return nil
}

type Log struct {
	Log uint `json:"log"`
}

func getData(e echo.Context) error {

	db := dbInit()

	var log model.Log

	l := db.First(&log, 1)

	err := l.Error
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, log)
	// return e.JSON(http.StatusOK, "data")
}

func createData(e echo.Context) error {
	log := new(Log)
	log.Log = 1

	if err := e.Bind(log); err != nil {
		return err
	}

	db := dbInit()

	l := db.Create(log)
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
