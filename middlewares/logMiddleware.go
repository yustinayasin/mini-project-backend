package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		//ini bisa custom mau dapat data apa
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}
