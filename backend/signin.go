package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func signin(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Printf("%+v", u)

	if inputPassword := getHashFromPassword(u.Password); inputPassword != u.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong password")
	}

	t, err := createToken(&SessionPayload{u.Id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, t)
}
