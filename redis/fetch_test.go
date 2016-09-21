package redis

import (
	"errors"
	"testing"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	mockConn := redigomock.NewConn()
	mockPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return mockConn, nil
		},
	}
	userService := &UserService{mockPool}

	t.Run("Found", found(userService, mockConn))
	t.Run("NotFound", notFound(userService, mockConn))
	t.Run("FetchError", fetchError(userService, mockConn))
}

func found(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		exp := rpimonitor.User{
			Login: "andre",
			Name:  "Andre Costa",
		}
		mockConn.Clear()
		mockConn.Command("GET", "user:andre:name").Expect(exp.Name)
		actual, err := userService.Fetch("andre")
		if assert.NoError(t, err) {
			assert.Equal(t, exp, actual)
		}
	}
}

func notFound(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		mockConn.Clear()
		mockConn.Command("GET", redigomock.NewAnyData()).Expect("")
		actual, err := userService.Fetch("not_andre")
		if assert.Error(t, err) {
			assert.Equal(t, rpimonitor.NotFound, err)
			assert.Equal(t, rpimonitor.User{}, actual)
		}
	}
}

func fetchError(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		mockConn.Clear()
		exp := errors.New("Connection error")
		mockConn.Command("GET", "user:andre:name").ExpectError(exp)
		_, err := userService.Fetch("andre")
		if assert.Error(t, err) {
			assert.Equal(t, exp, err)
		}
	}
}
