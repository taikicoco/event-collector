package handler

import (
	"github.com/labstack/echo/v4"
	"server/api/application"
)

func AssignRoutes(e *echo.Echo) {
	e.POST("/log/detail", application.CreateLogDetail)
	e.GET("/log/detail/:id", application.GetLogDetail)
	e.POST("/log/:id", application.CreateLog)
}
