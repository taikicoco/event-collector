package handler

import (
	"server/handler/request"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetLogName(e echo.Context) error {
	var req request.GetLogNameRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	res, err := h.logName.Get(ctx, req.ID)
	if err != nil {
		return err
	}
	return e.JSON(200, res)
}
