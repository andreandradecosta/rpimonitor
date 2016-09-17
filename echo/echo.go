package echo

import (
	"net/http"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type Server struct {
	StatusReader rpimonitor.StatusReader
}

func (s *Server) Start() {
	e := echo.New()
	e.SetLogLevel(log.ERROR)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} - ${method}, ${uri}, [${status}]\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "index")
	})
	e.GET("/status", s.status)
	e.Run(standard.New(":8080"))
}

func (s *Server) status(c echo.Context) error {
	status, err := s.StatusReader.ReadStatus()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, status)
}
