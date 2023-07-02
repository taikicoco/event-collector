package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

func GetCtx(ec echo.Context) context.Context {
	return ec.Request().Context()
}
