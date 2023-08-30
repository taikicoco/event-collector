package main

import (
	"fmt"
	"server/config"
	"server/handler"
	"server/repository"
	"server/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var cfg *config.ServerConfig
	cfg, err := config.LoadEnvConfig()
	if err != nil {
		panic(err)
	}

	db, err := repository.DBInit()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handlers := makeHandler(db)
	handler.SetupRoutes(e, handlers)
	if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		panic(err)
	}
}

func makeHandler(db *sqlx.DB) *handler.Handler {
	logName := usecase.NewLogName(db)
	logDetail := usecase.NewLogDetail(db)
	log := usecase.NewLog(db)

	return handler.NewHandler(logName, logDetail, log)
}
