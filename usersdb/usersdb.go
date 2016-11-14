package usersdb

import (
	"bufio"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/pkg/errors"
)

// UserService implements the rpimonitor.UserManager
type UserService struct {
	fileName string
}

// NewUserService configures the UsersDB to read user data from the file given.
func NewUserService(fileName string) *UserService {
	return &UserService{fileName: fileName}
}

// Authenticate returns a user if success to authenticate, and nil otherwise.
func (u *UserService) Authenticate(login, password string) (*rpimonitor.User, error) {
	file, err := os.Open(u.fileName)
	defer file.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "Authentication failed")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ":")
		l, n, p := row[0], row[1], row[2]
		if l == login {
			if bcrypt.CompareHashAndPassword([]byte(p), []byte(password)) == nil {
				return &rpimonitor.User{
					Login: l,
					Name:  n,
				}, nil
			}

		}
	}
	return nil, nil
}
