package application

import (
	"context"
	"server/api/domain/model"
	"server/api/repository"
)

type UpdateWebLogDataRequest struct {
	ID         uint `json:"id"`
	WebLogID   uint `json:"web_log_id"`
	Access     uint `json:"access"`
	Conversion uint `json:"conversion"`
}

func UpdateWebLogData(ctx context.Context, req *UpdateWebLogDataRequest) (*model.WebLogData, error) {
	webLogData := &model.WebLogData{
		ID:         req.ID,
		WebLogID:   req.WebLogID,
		Access:     req.Access,
		Conversion: req.Conversion,
	}
	wldr := repository.NewWebLogDataRepository()
	return wldr.Update(ctx, webLogData)
}
