package rpimonitor

import "errors"

type User struct {
	Login string
	Name  string
}

var NotFound = errors.New("User not found")

type UserManager interface {
	Authenticate(login, password string) (*User, error)
}
