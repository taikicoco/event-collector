package handler

import (
	"context"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func GetCtx(ec echo.Context) context.Context {
	return ec.Request().Context()
}

type Handler struct {
	logName usecase.LogName
}

func NewHandler(logName *usecase.LogName) *Handler {
	return &Handler{
		logName: *logName,
	}
}

func SetupRoutes(e *echo.Echo, h *Handler) {
	v1g := e.Group("/v1")
	{
		v1l := v1g.Group("/log")
		{
			v1ln := v1l.Group("/name")
			{
				v1ln.GET("/:id", h.GetLogName)
			}
		}
	}
}
