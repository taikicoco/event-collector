package handler

import (
	
	"net/http"
	"time"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
)

func AssignRoutes(e *echo.Echo) error {
	
	e.POST("/singup", singup)
	e.GET("/users:userId", getUser)
	e.PATCH("/users:userId", updateUser)
	e.POST("/close", deleteUser)

	return nil
}

type User struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`

	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func singup(e echo.Context) error {
	u := new(User)
		if err := e.Bind(u); err != nil {
			return err
	}

	db := dbInit()
	user := db.Create(u)
	return e.JSON(http.StatusOK, user)
}

func getUser(e echo.Context) error {
	return e.JSON(http.StatusOK, "getUsers")
}

func updateUser(e echo.Context) error {
	return e.JSON(http.StatusOK, "updateUser")
}

func deleteUser(e echo.Context) error {
	return e.JSON(http.StatusOK, "deleteUser")
}


func dbInit() *gorm.DB {
	dsn := "user:password@tcp(mysql)/db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

