package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type CreatingCharacterInfo struct {
	Name string `json:"name"`
	SkillValue
}

func createChar(c echo.Context) error {
	cci := &CreatingCharacterInfo{}
	if err := c.Bind(cci); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Debugf("%+v", cci)

	user, err := userFromContext(c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	csv, cvv, err := createCharacter(
		user.ID,
		cci.Name,
		"randomCompany",
		cci.Prog,
		cci.Testing,
		cci.Analyze,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	st := NewInternalState(csv, cvv)
	GetGlobalState().add(csv.ID, st)

	return c.JSON(http.StatusOK, from(st))
}
