package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gobase.com/base/pkg/api"
)

func initializeRoutes(e *echo.Echo, s *services) {
	//e.Use(middleware.BodyDump(logger))
	e.POST("/api/v0/users", api.RegisterUser(s.registering))
	e.Any("/api/test", test)
	//e.Use(authenticate)
}

func authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().URL.Query().Get("token")
		if len(token) > 0 {
			return next(c)
		}

		return echo.NewHTTPError(echo.ErrBadRequest.Code, "Failed authentication")
	}
}
func test(c echo.Context) error {
	return c.String(http.StatusOK, "Go base with echo is running!")
}
