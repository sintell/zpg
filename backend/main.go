package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	runMigrations()
}

func main() {
	e := echo.New()
	c := &ServerConfig{}
	err := ReadConfig(c, "backend/resources/config.json")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Printf("%+v", c)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(c.Addr()))
}
