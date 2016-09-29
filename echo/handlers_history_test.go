package echo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

type mockSampleFetcher struct {
	result []rpimonitor.Sample
	err    error
}

func (m *mockSampleFetcher) Query(start, end time.Time) ([]rpimonitor.Sample, error) {
	return m.result, m.err
}

func (m *mockSampleFetcher) resultAsJSON() string {
	res, _ := json.Marshal(m.result)
	return string(res)
}

func TestHistory(t *testing.T) {
	e := echo.New()
	mock := &mockSampleFetcher{}
	server := New("", WithSampleFetcher(mock))
	t.Run("No_Params", noParams(e, server))
	t.Run("Invalid_Date", invalidDate(e, server))
	t.Run("Return_Samples", returnSamples(e, mock, server))
	t.Run("Handle_Service_Error", handleServiceError(e, mock, server))
}

func noParams(e *echo.Echo, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, err := http.NewRequest(echo.GET, "/history", nil)
		if assert.NoError(t, err) {
			c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
			if assert.NoError(t, server.history(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Regexp(t, "error", rec.Body.String())
			}

		}
	}
}

func invalidDate(e *echo.Echo, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		q := make(url.Values)
		q.Set("start", "01-01-2001")
		q.Set("end", "01-01-2002")
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(echo.GET, "/history?"+q.Encode(), nil)
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		if assert.NoError(t, server.history(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Regexp(t, "error", rec.Body.String())
		}
	}
}

func returnSamples(e *echo.Echo, mock *mockSampleFetcher, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		q := make(url.Values)
		q.Set("start", "2016-09-20")
		q.Set("end", "2016-09-21")
		result := []rpimonitor.Sample{
			{
				LocalTime: time.Now(),
			},
		}
		mock.result = result
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(echo.GET, "/history?"+q.Encode(), nil)
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		if assert.NoError(t, server.history(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, mock.resultAsJSON(), rec.Body.String())
		}
	}
}

func handleServiceError(e *echo.Echo, mock *mockSampleFetcher, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		q := make(url.Values)
		q.Set("start", "2016-09-20")
		q.Set("end", "2016-09-21")
		mock.err = errors.New("Service Error")
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(echo.GET, "/history?"+q.Encode(), nil)
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		err := server.history(c)
		if assert.Error(t, err) {
			httpError, ok := err.(*echo.HTTPError)
			if assert.True(t, ok) {
				assert.Equal(t, http.StatusInternalServerError, httpError.Code)
				assert.Equal(t, mock.err.Error(), httpError.Error())
			}
		}
	}
}
