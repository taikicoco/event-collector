package admin

import (
	"context"
	"server/api/application"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Application application.ApplicationInterface
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func NewHandler(
	app application.ApplicationInterface,
) *AdminHandler {
	return &AdminHandler{
		Application: app,
	}
}

func (ah *AdminHandler) AssignRoutes(e *echo.Echo) {
	// v1g := e.Group("v1")
	// {
	// 	v1ag := v1g.Group("/admin")
	// 	{

	// 	}
	// }
}

func (ah *AdminHandler) GetCtx(ec echo.Context) context.Context {
	return ec.Request().Context()
}
