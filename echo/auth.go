package echo

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (s *Server) login(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	if ok, _ := s.UserManager.Authenticate(login, password); ok {
		user, _ := s.UserManager.Fetch(login)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"login": user.Login,
			"name":  user.Name,
		})
		tokenString, err := token.SignedString([]byte(s.JWTSigningKey))
		if err != nil {
			log.Println("Error signing token", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error during authentication.")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": tokenString,
		})
	}
	return echo.ErrUnauthorized
}
