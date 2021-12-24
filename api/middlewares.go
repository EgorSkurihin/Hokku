package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (srv *APIServer) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := srv.sessionStore.Get(c.Request(), "session")
		_, ok := session.Values["userId"]
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Wrong session")
		}
		return next(c)
	}
}
