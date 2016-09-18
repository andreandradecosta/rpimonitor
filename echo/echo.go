package echo

import (
	"fmt"
	"net/http"
	"time"

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
	e.GET("/", s.status)
	e.GET("/history", s.history)
	e.Run(standard.New("localhost:8080"))
}

func (s *Server) status(c echo.Context) error {
	status, err := s.StatusReader.ReadStatus()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, status)
}

func (s *Server) history(c echo.Context) error {
	start, errS := time.Parse("2006-01-02", c.QueryParam("start"))
	end, errE := time.Parse("2006-01-02", c.QueryParam("end"))
	if errS != nil || errE != nil {
		return c.String(http.StatusOK, "Invalid date range")
	}
	s.SampleFetcher.Query(start, end)
	return c.String(http.StatusOK, fmt.Sprintf("Result for %v - %v", start, end))
}
