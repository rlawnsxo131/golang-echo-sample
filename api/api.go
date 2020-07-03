package api

import (
	v1 "golang-echo-sample/api/v1"

	"github.com/labstack/echo/v4"
)

// ApplyRoutes ...
func ApplyRoutes(e *echo.Echo) {
	api := e.Group("/api")
	v1.ApplyRoutes(api)
}
