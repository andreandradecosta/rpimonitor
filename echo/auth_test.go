package echo

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

type mockUserManager struct {
	user *rpimonitor.User
	err  error
}

func (m *mockUserManager) Authenticate(login, password string) (*rpimonitor.User, error) {
	return m.user, m.err
}

func TestAuth(t *testing.T) {
	e := echo.New()
	mock := &mockUserManager{}
	server := New("", WithUserManager(mock))
	mock.user = &rpimonitor.User{
		Login: "user_login",
		Name:  "user_name",
	}
	t.Run("AuthOK", authOK(e, server))
	mock.user = nil
	t.Run("AuthFail", authFail(e, server))
	mock.err = errors.New("Service error")
	t.Run("AuthError", authError(e, server))
}

func authOK(e *echo.Echo, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		req, _ := http.NewRequest(echo.POST, "/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		if assert.NoError(t, server.login(c)) {
			assert.Regexp(t, "token", rec.Body.String())
		}
	}
}

func authFail(e *echo.Echo, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		req, _ := http.NewRequest(echo.POST, "/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		err := server.login(c)
		if assert.Error(t, err) {
			httpError, ok := err.(*echo.HTTPError)
			if assert.True(t, ok) {
				assert.Equal(t, http.StatusUnauthorized, httpError.Code)
			}
		}
	}
}

func authError(e *echo.Echo, server *Server) func(*testing.T) {
	return func(t *testing.T) {
		req, _ := http.NewRequest(echo.POST, "/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		err := server.login(c)
		if assert.Error(t, err) {
			httpError, ok := err.(*echo.HTTPError)
			if assert.True(t, ok) {
				assert.Equal(t, http.StatusInternalServerError, httpError.Code)
			}
		}
	}
}
