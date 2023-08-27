package main

import (
	"fmt"
	"server/handler"
	"server/repository"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := repository.DBInit()
	if err != nil {
		panic(err)
	}

	logName := usecase.NewLogName(db)
	logDetail := usecase.NewLogDetail(db)



	handlers := handler.NewHandler(logName, logDetail)

	e := echo.New()
	handler.SetupRoutes(e, handlers)
	if err := e.Start(fmt.Sprintf(":%s", "8080")); err != nil {
		panic(err)
	}
}
