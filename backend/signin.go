package main

import (
	"net/http"

	"github.com/go-pg/pg"

	"github.com/labstack/echo"
)

func signin(c echo.Context) error {
	userReq := &User{}
	if err := c.Bind(userReq); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	u := &User{}
	if err := GetDB().Model(u).Where("login = ?", userReq.Login).Select(); err != nil {
		if err == pg.ErrNoRows {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong login or password")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if inputPassword := getHashFromPassword(userReq.Password); inputPassword != u.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong login or password")
	}

	t, err := createToken(&SessionPayload{u.ID})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cs, err := u.getCharacter()
	if err != nil {
		if err == pg.ErrNoRows {
			return c.JSON(http.StatusOK, AuthResponse{Token: t})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, AuthResponse{Token: t, State: from(GetGlobalState().get(cs.ID))})
}
