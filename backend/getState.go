package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func getState(c echo.Context) error {
	user, err := userFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	cs, err := user.getCharacter()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, from(GetGlobalState().get(cs.ID)))
}
