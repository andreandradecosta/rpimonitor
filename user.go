package rpimonitor

// User represents a User in the application.
type User struct {
	Login string
	Name  string
}

// UserManager is the interface that must be implemented by the database service
// for fetching and authenticating users.
type UserManager interface {
	Authenticate(login, password string) (*User, error)
}
