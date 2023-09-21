package handler

import (
	"server/handler/request"
	"server/handler/response"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateEventLog(e echo.Context) error {
	var req request.CreateEventLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	// ctx := GetCtx(e)
	// log, err := h.log.Create(ctx, req.LogDetailID)
	// if err != nil {
	// 	return err
	// }

	res := response.CreateEventLogResponse{
		ID: 1,
	}

	return e.JSON(200, res)
}
