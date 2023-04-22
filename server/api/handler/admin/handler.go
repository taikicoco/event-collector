package admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AssignRoutes(e *echo.Echo) error {
	ag := e.Group("/admin")
	{
		ag.GET("/logs", getLogs)
	}
	return nil
}

func getLogs(e echo.Context) error {
	return e.JSON(http.StatusOK, "getLogs")
}
