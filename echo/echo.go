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
}

func (s *Server) Start() {
	e := echo.New()
	e.SetLogLevel(log.ERROR)
	e.SetDebug(true)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} - ${method}, ${uri}, [${status}]\n",
	}))
	e.Static("/about", "static")
	e.GET("/", s.status)
	e.GET("/history", s.history)
	e.Run(standard.New(":8080"))
}
