package handler

import (
	"server/handler/request"
	"server/handler/response"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateLog(e echo.Context) error {
	var req request.CreateLogRequest
	if err := e.Bind(&req); err != nil {
		return err
	}

	ctx := GetCtx(e)
	log, err := h.log.Create(ctx, req.LogDetailID)
	if err != nil {
		return err
	}

	res := response.CreateLogResponse{
		ID:          log.ID,
		LogDetailID: log.LogDetailID,
	}
	return e.JSON(200, res)
}
