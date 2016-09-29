package echo

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (s *Server) login(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	user, err := s.userManager.Authenticate(login, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "Error during auth process").Error())
	}
	if user == nil {
		return echo.ErrUnauthorized
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": user.Login,
		"name":  user.Name,
	})
	tokenString, err := token.SignedString([]byte(s.jwtSigningKey))
	if err != nil {
		log.Println("Error signing token", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error during authentication.")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
