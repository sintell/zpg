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
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	createCharacter(user.Id, cci.Name, "randomCompany", cci.Prog, cci.Testing, cci.Analyze)

	return c.JSON(http.StatusCreated, nil)
}
