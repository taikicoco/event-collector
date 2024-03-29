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
	logName   usecase.LogName
	logDetail usecase.LogDetail
	log       usecase.Log
}

func NewHandler(logName *usecase.LogName, logDetail *usecase.LogDetail,
	log *usecase.Log) *Handler {
	return &Handler{
		logName:   *logName,
		logDetail: *logDetail,
		log:       *log,
	}
}

func SetupRoutes(e *echo.Echo, h *Handler) {
	v1g := e.Group("/v1")
	{
		// v1l := v1g.Group("/log")
		// {
		// 	v1ln := v1l.Group("/name")
		// 	{
		// 		v1ln.GET("/:id", h.GetLogName)
		// 	}
		// 	v1l.GET("/:id", h.GetLogDetail)
		// 	v1l.POST("/:id", h.CreateLog)
		// }

		v1e := v1g.Group("/event_groups")
		{
			v1e.POST("/:id", h.mock)

		}

		v1el := v1e.Group("/event")
		{
			v1el.POST("/:id", h.mock)
			v1el := v1e.Group("/log")
			{
				v1el.POST("/:id", h.CreateEventLog)
			}
		}
	}
}

func (h *Handler) mock(e echo.Context) error {
	return nil
}
