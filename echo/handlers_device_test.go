package echo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

type mockDevice struct {
	sample *rpimonitor.Sample
	status *rpimonitor.Status
	err    error
}

func (m *mockDevice) ReadSample() (*rpimonitor.Sample, error) {
	return m.sample, m.err
}

func (m *mockDevice) ReadStatus() (*rpimonitor.Status, error) {
	return m.status, m.err
}

func TestDevice(t *testing.T) {
	e := echo.New()
	mock := &mockDevice{}
	server := New("", WithDevice(mock))
	mock.status = &rpimonitor.Status{}
	mock.sample = &rpimonitor.Sample{}
	t.Run("StatusOk", readOK(e, mock.status, server.status))
	t.Run("SnapshotOk", readOK(e, mock.sample, server.snapshot))
	mock.err = errors.New("Device error")
	t.Run("StatusError", readError(e, mock.err, server.status))
	t.Run("SnapshotError", readError(e, mock.err, server.snapshot))
}

func readOK(e *echo.Echo, exp interface{}, f echo.HandlerFunc) func(*testing.T) {
	return func(t *testing.T) {
		expJSON, _ := json.Marshal(exp)
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(echo.GET, "/api/service", nil)
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		if assert.NoError(t, f(c)) {
			assert.JSONEq(t, string(expJSON), rec.Body.String())
		}
	}
}

func readError(e *echo.Echo, expErr error, f echo.HandlerFunc) func(*testing.T) {
	return func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(echo.GET, "/api/service", nil)
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		err := f(c)
		if assert.Error(t, err) {
			httpError, ok := err.(*echo.HTTPError)
			if assert.True(t, ok) {
				assert.Equal(t, http.StatusInternalServerError, httpError.Code)
				assert.Equal(t, expErr.Error(), httpError.Message)
			}
		}
	}
}
