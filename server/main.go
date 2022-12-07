package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	// start server
	e.Logger.Fatal(e.Start(":1222"))
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
