package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Server struct {
	*echo.Echo
	conf *ServerConfig
}

var server *Server

func initServer() *Server {
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

	return &Server{e, c}
}

func GetServer() *Server {
	if server != nil {
		return server
	}

	return initServer()
}

func (s *Server) Run() {
	s.Logger.Fatal(s.Start(s.conf.Addr()))
}
