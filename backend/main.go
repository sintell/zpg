package main

import (
	"net/http"

	"github.com/go-pg/pg"

	"github.com/labstack/echo"
)

type User struct {
	Login    string `sql:"login"`
	Password string `sql:"password"`
}

func init() {
	db := pg.Connect(&pg.Options{Addr: "localhost:5432"})

	db.CreateTable(&User{}, nil)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
