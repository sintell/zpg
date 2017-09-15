package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func createChar(c echo.Context) error {
	type CreatingCharacterInfo struct {
		Name string `json:"name"`
		SkillValue
	}

	cci := &CreatingCharacterInfo{}
	if err := c.Bind(cci); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := userFromContext(c)
	if err != nil {
		c.Logger().Warn(err.Error())
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

	GetGlobalState().add(csv.ID, NewInternalState(csv, cvv))

	return c.JSON(http.StatusCreated, nil)
}
