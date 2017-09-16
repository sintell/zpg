package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func rating(c echo.Context) error {
	if rp, err := getRatingByExp(); err == nil {
		return c.JSON(http.StatusOK, rp)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
}
