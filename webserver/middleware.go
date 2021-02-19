package webserver

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var (
	ErrNoAuth           = errors.New("Authorization header is missing")
	ErrNoToken          = errors.New("Token is missing")
	ErrInvalidTokenType = errors.New("Authorization Bearer is required")
	ErrJWTMissing       = errors.New("JWT token is missing")
	ErrJWTFormat        = errors.New("Wrong JWT token format")
)

func Authorization(s *Webserver) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token, err := getToken(c)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
				return s.secret, nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			info := t.Claims.(jwt.MapClaims)

			c.Set("info", info)

			return next(c)
		}
	}
}

func getToken(c echo.Context) (string, error) {
	auth := c.Request().Header.Get("Authorization")

	if len(auth) >= 1 && strings.HasPrefix(auth, "Bearer") {
		splitToken := strings.Split(auth, "Bearer ")
		if len(splitToken[1]) < 1 {
			return "", ErrJWTFormat
		}

		return splitToken[1], nil
	}

	return "", ErrJWTMissing
}
