package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

func GetCtx(ec echo.Context) context.Context {
	return ec.Request().Context()
}

type Handlers struct {
	WebLogHandler     WebLogHandlerInterface
	WebLogDataHandler WebLogDataHandlerInterface
}

func SetupRoutes(e *echo.Echo, handlers *Handlers) {
	v1g := e.Group("/v1")
	{
		v1wl := v1g.Group("/web_log")
		{
			// GET
			v1wl.GET("", handlers.WebLogHandler.GetWebLogs)
			v1wl.GET("/:id", handlers.WebLogHandler.GetWebLog)

			// POST
			v1wl.POST("", handlers.WebLogHandler.CreateWebLog)

			// PUT
			v1wl.PUT("/:id", handlers.WebLogHandler.UpdateWebLog)

			v1wld := v1wl.Group("/data")
			{
				// POST
				v1wld.PUT("/:id", handlers.WebLogDataHandler.UpdateWebLogData)
			}
		}
	}
}
