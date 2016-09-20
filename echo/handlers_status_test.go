package echo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

type mockStatusReader struct {
	result rpimonitor.Status
	err    error
}

func (m *mockStatusReader) ReadStatus() (rpimonitor.Status, error) {
	return m.result, m.err
}

func (m *mockStatusReader) resultAsJSON() string {
	res, _ := json.Marshal(m.result)
	return string(res)
}

func TestStatus(t *testing.T) {
	e := echo.New()
	mock := &mockStatusReader{}
	server := &Server{
		StatusReader: mock,
	}
	t.Run("Returns_Response", statusSuccess(e, mock, server))
	t.Run("Return_Error", statusError(e, mock, server))
}

func statusSuccess(e *echo.Echo, mock *mockStatusReader, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		mock := &mockStatusReader{
			result: rpimonitor.Status{LocalTime: time.Now()},
		}
		server := &Server{
			StatusReader: mock,
		}
		if assert.NoError(t, server.status(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, mock.resultAsJSON(), rec.Body.String())
		}
	}
}

func statusError(e *echo.Echo, mock *mockStatusReader, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		mock.err = errors.New("Device Error")
		err := server.status(c)
		if assert.Error(t, err) {
			httpError, ok := err.(*echo.HTTPError)
			if assert.True(t, ok) {
				assert.Equal(t, http.StatusInternalServerError, httpError.Code)
				assert.Equal(t, mock.err.Error(), httpError.Message)
			}
		}
	}
}
