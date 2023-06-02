package seed

import (
	"fmt"
	"server/api/domain/model"
	"time"

	"gorm.io/gorm"
)

func AccessLogSeed(db *gorm.DB) error {

	var recodeCount int64
	db.Table("access_logs").Count(&recodeCount)
	fmt.Println(recodeCount)

	if recodeCount > 0 {
		fmt.Println("AccessLogSeed Skip")
		return nil
	}

	accessLog := model.AccessLog{
		AccessLogDetailID: 1,
		Access:            10, Conversion: 2,
		CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err := db.Create(&accessLog).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println("AccessLogSeed done")
	return nil
}
