package main

import (
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

type Server struct {
	*echo.Echo
	conf *ServerConfig
}

var server *Server

func initServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	c := &ServerConfig{}
	err := ReadConfig(c, *configPath)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Printf("%+v", c)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/signup", signup)
	e.POST("/signin", signin)
	e.GET("/ratings", signup)

	withAuth := e.Group("/game", middleware.JWT([]byte("secret")))

	withAuth.GET("/state", signup)
	withAuth.POST("/create_char", createChar)
	withAuth.POST("/rest", rest)
	withAuth.POST("/test", testToken)

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
