package redis

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/andreandradecosta/rpimonitor"
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
		exp := &rpimonitor.User{
			Login: "andre",
			Name:  "Andre Costa",
		}
		pass := "password"
		hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		servRes := map[string]string{
			"name": exp.Name,
			"hash": string(hash),
		}
		mockConn.Clear()
		mockConn.Command("HGETALL", "user:andre").ExpectMap(servRes)
		act, err := userService.Authenticate("andre", pass)
		if assert.NoError(t, err) {
			assert.Equal(t, exp, act)
		}
	}
}

func authFail(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		mockConn.Clear()
		mockConn.Command("HGETALL", redigomock.NewAnyData()).ExpectMap(map[string]string{})
		act, err := userService.Authenticate("not_andre", "his pass")
		if assert.NoError(t, err) {
			assert.Nil(t, act)
		}
	}
}

func authError(userService *UserService, mockConn *redigomock.Conn) func(*testing.T) {
	return func(t *testing.T) {
		mockConn.Clear()
		exp := errors.New("Connection error")
		mockConn.Command("HGETALL", redigomock.NewAnyData()).ExpectError(exp)
		_, err := userService.Authenticate("andre", "password")
		if assert.Error(t, err) {
			assert.Equal(t, exp, errors.Cause(err))
		}
	}
}
