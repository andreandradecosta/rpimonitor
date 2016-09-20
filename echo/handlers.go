package echo

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type handleError struct {
	Msg string `json:"Error"`
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
		return c.JSON(http.StatusOK, &handleError{"Invalid date range"})
	}
	res, err := s.SampleFetcher.Query(start, end)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
