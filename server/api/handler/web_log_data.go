package handler

import (
	"net/http"
	"server/api/application"
	"server/api/handler/request"

	"github.com/labstack/echo/v4"
)

func UpdateWebLogData(e echo.Context) error {
	var req request.UpdateWebLogDataRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := application.UpdateWebLogData(ctx, &application.UpdateWebLogDataRequest{
		ID:         req.ID,
		WebLogID:   req.WebLogID,
		Access:     req.Access,
		Conversion: req.Conversion,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}