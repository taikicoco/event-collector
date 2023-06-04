package seed

import (
	"fmt"
	"server/api/domain/model"
	"time"

	"gorm.io/gorm"
)

func LogSeed(db *gorm.DB) error {

	var recodeCount int64
	db.Table("logs").Count(&recodeCount)
	fmt.Println(recodeCount)

	if recodeCount > 0 {
		fmt.Println("LogSeed Skip")
		return nil
	}

	Log := model.Log{
		LogDetailID: 1,
		Access:      10, Conversion: 2,
		CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err := db.Create(&Log).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println("LogSeed done")
	return nil
}
