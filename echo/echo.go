package echo

import (
	"github.com/andreandradecosta/rpimonitor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type Server struct {
	StatusReader  rpimonitor.StatusReader
	SampleFetcher rpimonitor.SampleFetcher
	SampleReader  rpimonitor.SampleReader
	UserManager   rpimonitor.UserManager
	JWTSigningKey string
	Debug         bool
}

func (s *Server) Start() {
	e := echo.New()
	e.SetLogLevel(log.ERROR)
	e.SetDebug(s.Debug)
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Secure())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} - ${method}, ${uri}, [${status}]\n",
	}))

	e.Static("/", "static")
	e.POST("/login", s.login)

	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(s.JWTSigningKey),
		SigningMethod: "HS256",
	}))
	r.GET("/status", s.status)
	r.GET("/history", s.history)
	r.GET("/snapshot", s.snapshot)
	e.Run(standard.WithTLS(
		":8443",
		"cert.pem",
		"key.pem",
	))
}
