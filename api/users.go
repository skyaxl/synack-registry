package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/skyaxl/synack-registry/pkg/apierrors"
	"github.com/skyaxl/synack-registry/pkg/users/userscontracts"
)

type UsersHandler struct {
	service userscontracts.UserService
}

func (h *UsersHandler) Bind(router echo.Echo) {
	router.GET("/users/:username", h.GET)
	router.POST("/users", h.POST)
	router.PUT("/users/:username", h.PUT)
	router.GET("/users/:username", h.DELETE)

}

//GET User
func (h *UsersHandler) GET(c echo.Context) error {
	username := c.Param("username")
	ctx := c.Request().Context()
	user, err := h.service.Get(ctx, username)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, user)
	return nil
}

//POST User
func (h *UsersHandler) POST(c echo.Context) error {
	ctx := c.Request().Context()
	user := userscontracts.User{}
	if err := c.Bind(&user); err != nil {
		return errors.WithMessage(apierrors.ErrBadRequest, err.Error())
	}

	if err := h.service.Save(ctx, user); err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

//PUT
func (h *UsersHandler) PUT(c echo.Context) error {
	ctx := c.Request().Context()
	username := c.Param("username")
	user := userscontracts.User{}
	if err := c.Bind(&user); err != nil || username != user.Username {
		return errors.WithMessage(apierrors.ErrBadRequest, err.Error())
	}

	if err := h.service.Update(ctx, user); err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

//DELETE User
func (h *UsersHandler) DELETE(c echo.Context) error {
	username := c.Param("username")
	ctx := c.Request().Context()
	if err := h.service.Delete(ctx, username); err != nil {
		return err
	}
	c.String(http.StatusOK, "{}")
	return nil
}
