package api

import (
	"github.com/labstack/echo"
	"github.com/skyaxl/synack-registry/pkg/apierrors"
	"github.com/skyaxl/synack-registry/pkg/users/userscontracts"
)

type AuthHandler struct {
	user userscontracts.UserService
}

func (ah *AuthHandler) Middleware(username, password string, c echo.Context) (bool, error) {
	ctx := c.Request().Context()
	authenticatedUser, err := ah.user.Autenticate(ctx, userscontracts.AuthenticationRequest{
		Username: username,
		Password: password,
	})

	if err == nil {
		c.Set("authenticated_user", authenticatedUser)
		return true, nil
	}

	var apiError apierrors.ApiError
	var ok bool
	if apiError, ok = err.(apierrors.ApiError); !ok {
		apiError = apierrors.ErrInternalServerError
	}
	return true, apiError
}
