package middleware

import (
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, stm := utils.Authorize(c)
		if stm {
			return next(c)
		}
		return c.Redirect(302, "/")
	}
}

func NotAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, stm := utils.Authorize(c)
		if !stm {
			return next(c)
		}
		return c.Redirect(302, "/")
	}
}
