package main

import (
	"crypto/sha1"
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo"
)

func signup(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Printf("%+v", u)

	hash := sha1.New()
	hash.Write([]byte(u.Password))
	u.Password = base64.URLEncoding.EncodeToString(hash.Sum(nil))

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

	return c.JSON(http.StatusOK, t)
}
