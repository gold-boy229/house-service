package middleware

import (
	"house-store/internal/consts"
	"house-store/internal/utilities/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthorizedModeratorsOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := getAuthToken(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		paramsFromToken, err := auth.ValidateAuthorizedToken(tokenString)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		userRole := paramsFromToken.UserRole
		userUUID := paramsFromToken.UserUUID

		// validate extracted user role
		if !auth.IsModerator(userRole) {
			return c.String(http.StatusForbidden, "Forbidden. You don't have enough rights for this action")
		}

		c.Set(consts.ECHO_CONTEXT_USER_ROLE_KEY, userRole)
		c.Set(consts.ECHO_CONTEXT_USER_UUID_KEY, userUUID)
		return next(c)
	}
}
