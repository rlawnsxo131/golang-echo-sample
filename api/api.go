package api

import (
	v1 "golang-echo-sample/api/v1"

	"github.com/labstack/echo/v4"
)

func ApplyRoutes(e *echo.Group) {
	api := e.Group("/api")
	v1.ApplyRoutes(api)
}
