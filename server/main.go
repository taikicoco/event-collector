package main

import (
	"fmt"
	"server/api/handler"
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

	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err = handler.AssignRoutes(e); err != nil {
		panic(err)
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
