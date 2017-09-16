package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func rest (c echo.Context) error {
	user, err := userFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	charStat, err := user.getCharacter()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	GetGlobalState().get(charStat.ID).CharVarValue.rest()

}
