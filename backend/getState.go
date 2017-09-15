package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func getState(c echo.Context) error {
	user, err := userFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, from(GetGlobalState().get(user.Id)))
}
