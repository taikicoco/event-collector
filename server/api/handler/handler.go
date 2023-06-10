package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

func AssignRoutes(e *echo.Echo) {
	v1g := e.Group("/v1")
	{
		v1wl := v1g.Group("/web_log")
		{
			// GET
			v1wl.GET("", GetWebLogs)
			v1wl.GET("/:id", GetWebLog)

			// POST
			v1wl.POST("", CreateWebLog)

			// PUT
			v1wl.PUT("/:id", UpdateWebLog)

			v1wld := v1wl.Group("/data")
			{
				// POST
				v1wld.PUT("/:id", UpdateWebLogData)
			}
		}
	}
}

func GetCtx(ec echo.Context) context.Context {
	return ec.Request().Context()
}
