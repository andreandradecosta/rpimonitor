package redis

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	mockConn := redigomock.NewConn()
	mockPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return mockConn, nil
		},
	}
	userService := &UserService{mockPool}
	t.Run("Auth_OK", authOk(userService, mockConn))
	t.Run("Auth_Fail", authFail(userService, mockConn))
	t.Run("Auth_Error", authError(userService, mockConn))
}

func authOk(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		pass := "password"
		hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		mockConn.Clear()
		mockConn.Command("GET", "user:andre:hash").Expect(hash)
		aut, err := userService.Authenticate("andre", pass)
		if assert.NoError(t, err) {
			assert.True(t, aut)
		}
	}
}

func authFail(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		mockConn.Clear()
		mockConn.Command("GET", redigomock.NewAnyData()).Expect("")
		aut, err := userService.Authenticate("not_andre", "his pass")
		if assert.NoError(t, err) {
			assert.False(t, aut)
		}
	}
}

func authError(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		mockConn.Clear()
		exp := errors.New("Connection error")
		mockConn.Command("GET", "user:andre:hash").ExpectError(exp)
		_, err := userService.Authenticate("andre", "password")
		if assert.Error(t, err) {
			assert.Equal(t, exp, errors.Cause(err))
		}
	}
}
