package middleware

import (
	"errors"
	"fmt"
	"house-store/internal/consts"
	"house-store/internal/utilities/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ModeratorsOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole, err := getUserRole(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		// validate extracted user role
		if !auth.IsModerator(userRole) {
			return c.String(http.StatusForbidden, "Forbidden. You don't have enough rights for this action")
		}

		c.Set(consts.ECHO_CONTEXT_USER_ROLE_KEY, userRole)
		return next(c)
	}
}

func AuthOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole, err := getUserRole(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		// validate extracted user role
		if !auth.IsClient(userRole) && !auth.IsModerator(userRole) {
			return c.String(http.StatusForbidden, "Forbidden. You don't have enough rights for this action")
		}

		// set user role for next functions
		c.Set(consts.ECHO_CONTEXT_USER_ROLE_KEY, userRole)

		return next(c)
	}
}

func getUserRole(c echo.Context) (string, error) {
	token, err := getAuthToken(c)
	if err != nil {
		return "", err
	}

	userRole, err := extractUserRole(token)
	if err != nil {
		return "", err
	}
	return userRole, nil
}

func getAuthToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no Authorization header was provided")
	}

	// Expected format: "Bearer <TOKEN>"
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid Authorization header format")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token, nil
}

func extractUserRole(token string) (string, error) {
	userRole, err := auth.ValidateToken(token)
	if err != nil {
		return "", fmt.Errorf("invalid Authorization token: %v", err.Error())
	}
	return userRole, nil
}
