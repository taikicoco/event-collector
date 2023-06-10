package handler

import (
	"net/http"
	"server/api/application"
	"server/api/handler/request"

	"github.com/labstack/echo/v4"
)

func GetWebLogs(e echo.Context) error {
	ctx := GetCtx(e)

	res, err := application.GetWebLogs(ctx)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}

func GetWebLog(e echo.Context) error {
	var req request.GetWebLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := application.GetWebLog(ctx, req.ID)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}

func CreateWebLog(e echo.Context) error {
	var req request.CreateWebLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := application.CreateWebLog(ctx, &application.CreateWebLogRequest{
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}

func UpdateWebLog(e echo.Context) error {
	var req request.UpdateWebLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := application.UpdateWebLog(ctx, &application.UpdateWebLogRequest{
		ID:      req.ID,
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}
