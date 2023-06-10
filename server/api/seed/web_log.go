package seed

import (
	"fmt"
	"server/api/domain/model"
	"time"

	"gorm.io/gorm"
)

func WebLogSeed(db *gorm.DB) error {
	var recodeCount int64
	db.Table("web_logs").Count(&recodeCount)
	fmt.Println("web_log recode = ", recodeCount)

	if recodeCount > 0 {
		fmt.Println("WebLogSeed Skip")
		return nil
	}

	WebLog := model.WebLog{
		Name: "detail_name", Version: 1, PageUrl: "www.exsampl.com",
		CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err := db.Create(&WebLog).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println("WebLogSeed done")
	return nil
}
