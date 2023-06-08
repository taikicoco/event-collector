package seed

import (
	"fmt"
	"server/api/domain/model"
	"time"

	"gorm.io/gorm"
)

func WebLogDataSeed(db *gorm.DB) error {
	var recodeCount int64
	db.Table("web_log_data").Count(&recodeCount)
	fmt.Println("web_log_data recode = ", recodeCount)

	if recodeCount > 0 {
		fmt.Println("WebLogDataSeed Skip")
		return nil
	}

	webLogData := model.WebLogData{
		WebLogID: 1, Access: 1, Conversion: 1,
		CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err := db.Create(&webLogData).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println("WebLogDataSeed done")
	return nil
}
