package handler

import (
	"net/http"
	"server/api/application"
	"github.com/labstack/echo/v4"
)


func CreateLog(e echo.Context) error {
	var req application.CreateLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}
	ctx := GetCtx(e)

	res, err := application.CreateLog(ctx, &application.CreateLogRequest{
		LogDetailID: req.LogDetailID,
		Access:      req.Access,
		Conversion:  req.Conversion,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}