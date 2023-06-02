package seed

import (
	"fmt"
	"server/api/domain/model"
	"time"

	"gorm.io/gorm"
)

func AccessLogDetailSeed(db *gorm.DB) error {
	var recodeCount int64
	db.Table("access_log_details").Count(&recodeCount)
	fmt.Println(recodeCount)

	if recodeCount > 0 {
		fmt.Println("AccessLogDetailSeed Skip")
		return nil
	}

	accessLog := model.AccessLogDetail{
		Name: "detail_name", Version: 1, PageUrl: "www.exsampl.com",
		CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err := db.Create(&accessLog).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println("AccessLogSeedDetail done")
	return nil
}
