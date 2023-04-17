package main

import (
	"fmt"
	"net/http"

	"server/config"
	"server/api/application"
	"server/api/infrastracture/persistence"
	admin_handler "server/api/handler/admin"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	cfg *config.ServerConfig
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	// start server
	e.Logger.Fatal(e.Start(":1222"))

	if err = assignRoutes(e, cfg); err != nil {
		panic(err)
	}

	if err = e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		panic(err)
	}
}

func assignRoutes(e *echo.Echo, cfg *config.ServerConfig) error {
	// Persistence
	ri, err := persistence.Connect(cfg)
	if err != nil {
		return err
	}
	
	// Application
	bdl := &application.ApplicationBundle{
		ServerConfig:             cfg,
		Repository:               ri,
	}
	app := application.NewApplication(bdl)

	// Admin Handler
	ah := admin_handler.NewHandler(app)
	ah.AssignRoutes(e)

	return nil
}

func hello(ec echo.Context) error {
	db := dbInit()
	db.AutoMigrate(&Demo{})
	insert(db)

	return ec.JSON(http.StatusCreated, "ok")
}

func dbInit() *gorm.DB {
	dsn := "user:password@tcp(mysql)/db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func insert(db *gorm.DB) {
	demo := Demo{
		Name: "demo_name",
		ID:   1,
	}
	result := db.Create(&demo)

	fmt.Println("count:", result.RowsAffected)
}

type Demo struct {
	gorm.Model

	Name string
	ID   int
}
