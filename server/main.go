package main

import (
	"fmt"
	"server/api/application"
	"server/api/handler"
	"server/api/infrastructure"
	"server/api/seed"
	"server/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	err error
	cfg *config.ServerConfig
)

func main() {
	if cfg, err = config.LoadEnvConfig(); err != nil {
		panic(err)
	}

	// seeder
	seedInit(cfg)

	// echo instance
	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers := initHandlers()
	handler.SetupRoutes(e, handlers)

	fmt.Println(cfg.Port, "port")
	if err = e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		panic(err)
	}
}

func seedInit(cfg *config.ServerConfig) {
	if cfg.Environment == "development" {
		seed.Seed()
	}
}

func initHandlers() *handler.Handlers {
	webLogRepo := infrastructure.NewWebLogRepository(config.NewDB())
	webLogApp := application.NewWebLogApplication(webLogRepo)
	webLogHandler := handler.NewWebLogHandler(webLogApp)

	webLogDataRepo := infrastructure.NewWebLogDataRepository(config.NewDB())
	webLogDataApp := application.NewWebLogDataApplication(webLogDataRepo)
	webLogDataHandler := handler.NewWebLogDataHandler(webLogDataApp)

	return &handler.Handlers{
		WebLogHandler:     webLogHandler,
		WebLogDataHandler: webLogDataHandler,
	}
}
