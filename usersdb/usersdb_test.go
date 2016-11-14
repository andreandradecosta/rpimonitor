package usersdb

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	userService := setup(t)
	defer os.Remove(userService.fileName)
	t.Run("Users File Missing", userFileMissing(NewUserService("non existed")))
	t.Run("User Not Found", userNotFound(userService))
	t.Run("Password Mismatch", passwordMismatch(userService))
	t.Run("Auth OK", authOK(userService))
}

func setup(t *testing.T) *UserService {

	users := `user1:User Name 1:$2a$06$9tV.C.JQ3e/uLldLKw6AG.Oq9u2pQOY4prDtEA2ub6FAHPUy5eG0.
user2:User Name 2:$2a$06$rRPyoBBWY3QsMNaft5.mhOirujXtQylQYkH.8CFdD8ZoluQ8dUNCS`

	tmpfile, err := ioutil.TempFile(".", "users_file_")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write([]byte(users)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	return &UserService{tmpfile.Name()}
}

func userFileMissing(u *UserService) func(*testing.T) {
	return func(t *testing.T) {
		_, err := u.Authenticate("login", "password")
		assert.Error(t, err)
	}
}

func userNotFound(u *UserService) func(*testing.T) {
	return func(t *testing.T) {
		actual, err := u.Authenticate("user3", "passwd1")
		if assert.NoError(t, err) {
			assert.Nil(t, actual)
		}
	}
}

func passwordMismatch(u *UserService) func(*testing.T) {
	return func(t *testing.T) {
		actual, err := u.Authenticate("user2", "passwd1")
		if assert.NoError(t, err) {
			assert.Nil(t, actual)
		}
	}
}

func authOK(u *UserService) func(*testing.T) {
	return func(t *testing.T) {
		exp := &rpimonitor.User{
			Login: "user1",
			Name:  "User Name 1",
		}
		actual, err := u.Authenticate("user1", "passwd1")
		if assert.NoError(t, err) && assert.NotNil(t, actual) {
			assert.Equal(t, exp, actual)
		}
	}
}
