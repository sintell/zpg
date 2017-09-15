package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func signin(c echo.Context) error {
	userReq := &User{}
	if err := c.Bind(userReq); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	u := &User{}
	GetDB().Model(u).Where("login=?", userReq.Login).Select()
	if inputPassword := getHashFromPassword(userReq.Password); inputPassword != u.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong password")
	}

	t, err := createToken(&SessionPayload{u.ID})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cs, err := u.getCharacter()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, AuthResponse{Token: t, State: from(GetGlobalState().get(cs.ID))})
}
