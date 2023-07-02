package handler

import (
	"net/http"
	"server/api/application"
	"server/api/handler/request"

	"github.com/labstack/echo/v4"
)

type WebLogHandlerInterface interface {
	GetWebLogs(echo.Context) error
	GetWebLog(echo.Context) error
	CreateWebLog(echo.Context) error
	UpdateWebLog(echo.Context) error
}

type webLogHandler struct {
	webLogApp application.WebLogApplicationInterface
}

func NewWebLogHandler(webLogApp application.WebLogApplicationInterface) WebLogHandlerInterface {
	return &webLogHandler{webLogApp: webLogApp}
}

func (wlh webLogHandler) GetWebLogs(e echo.Context) error {
	ctx := GetCtx(e)

	res, err := wlh.webLogApp.GetWebLogs(ctx)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}

func (wlh webLogHandler) GetWebLog(e echo.Context) error {
	var req request.GetWebLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := wlh.webLogApp.GetWebLog(ctx, req.ID)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}

func (wlh webLogHandler) CreateWebLog(e echo.Context) error {
	var req request.CreateWebLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := wlh.webLogApp.CreateWebLog(ctx, &application.CreateWebLogRequest{
		Name:    req.Name,
		Version: req.Version,
		PageUrl: req.PageUrl,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, res)
}

func (wlh webLogHandler) UpdateWebLog(e echo.Context) error {
	var req request.UpdateWebLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := wlh.webLogApp.UpdateWebLog(ctx, &application.UpdateWebLogRequest{
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
