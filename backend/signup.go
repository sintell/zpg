package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type AuthResponse struct {
	Token string `json:"token,omitempty"`
}

func signup(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Printf("%+v", u)

	u.Password = getHashFromPassword(u.Password)

	err := NewUser(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var usr []User
	GetDB().Model(&usr).Select()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Printf("%+v", usr)

	t, err := createToken(&SessionPayload{u.Id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, AuthResponse{Token: t})
}
