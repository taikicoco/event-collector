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
	seed_init(cfg)

	webLogRepo := infrastructure.NewWebLogRepository(config.NewDB())
	webLogApp := application.NewWebLogApplication(webLogRepo)
	webLogHandler := handler.NewWebLogHandler(webLogApp)

	webLogDataRepo := infrastructure.NewWebLogDataRepository(config.NewDB())
	webLogDataApp := application.NewWebLogDataApplication(webLogDataRepo)
	webLogDataHandler := handler.NewWebLogDataHandler(webLogDataApp)

	// echo instance
	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// routing
	v1g := e.Group("/v1")
	{
		v1wl := v1g.Group("/web_log")
		{
			// GET
			v1wl.GET("", webLogHandler.GetWebLogs)
			v1wl.GET("/:id", webLogHandler.GetWebLog)

			// POST
			v1wl.POST("", webLogHandler.CreateWebLog)

			// PUT
			v1wl.PUT("/:id", webLogHandler.UpdateWebLog)

			v1wld := v1wl.Group("/data")
			{
				// POST
				v1wld.PUT("/:id", webLogDataHandler.UpdateWebLogData)
			}
		}
	}

	fmt.Println(cfg.Port, "port")
	if err = e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		panic(err)
	}
}

func seed_init(cfg *config.ServerConfig) {
	if cfg.Environment == "development" {
		seed.Seed()
	}
}
