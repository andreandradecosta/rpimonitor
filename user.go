package rpimonitor

import "errors"

type User struct {
	Login string
	Name  string
}

var NotFound = errors.New("User not found")

type UserManager interface {
	Fetch(login string) (User, error)
	Authenticate(login, password string) (bool, error)
}
