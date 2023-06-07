package handler

import (
	"context"
	"github.com/labstack/echo/v4"
)

func AssignRoutes(e *echo.Echo) {
	e.POST("/log/detail", CreateLogDetail)
	e.GET("/log/detail/:id", GetLogDetail)
	e.POST("/log/:id", CreateLog)
}

func GetCtx(ec echo.Context) context.Context {
	return ec.Request().Context()
}
