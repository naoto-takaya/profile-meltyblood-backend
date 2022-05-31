package configs

import (
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Config
}

func CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{}
		cc.Context = c
		return next(cc)
	}
}

func ConfigMiddleware(config Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*CustomContext)
			cc.Config = config
			return next(cc)
		}
	}
}
