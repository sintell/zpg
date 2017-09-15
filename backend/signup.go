package main

import (
	"github.com/labstack/echo"
)

func init() {
	GetServer().GET("/signup", signup)
}

func signup(c echo.Context) error {
	return nil
}
