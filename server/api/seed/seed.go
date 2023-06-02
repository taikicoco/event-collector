package seed

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbInit() *gorm.DB {
	dsn := "user:password@tcp(mysql)/db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Seed() {
	db := dbInit()

	if err := AccessLogSeed(db); err != nil {
		fmt.Printf("%+v", err)
	}

	if err := AccessLogDetailSeed(db); err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println("Seed done")
}
