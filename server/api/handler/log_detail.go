package handler

import (
	"net/http"
	"server/api/application"
	
	"github.com/labstack/echo/v4"
)

func CreateLogDetail(e echo.Context) error {
	var req application.CreateLogDetailRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := application.CreateLogDetail(ctx, &application.CreateLogDetailRequest{
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}

func GetLogDetail(e echo.Context) error {
	var req application.GetLogDetailRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := application.GetLogDetail(ctx, req.ID)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}
