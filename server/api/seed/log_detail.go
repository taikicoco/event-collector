package seed

import (
	"fmt"
	"server/api/domain/model"
	"time"

	"gorm.io/gorm"
)

func LogDetailSeed(db *gorm.DB) error {
	var recodeCount int64
	db.Table("log_details").Count(&recodeCount)
	fmt.Println(recodeCount)

	if recodeCount > 0 {
		fmt.Println("LogDetailSeed Skip")
		return nil
	}

	Log := model.LogDetail{
		Name: "detail_name", Version: 1, PageUrl: "www.exsampl.com",
		CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err := db.Create(&Log).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println("LogSeedDetail done")
	return nil
}
