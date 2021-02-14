package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/skyaxl/synack-registry/pkg/apierrors"
	"github.com/skyaxl/synack-registry/pkg/regs/regscontract"
	"github.com/skyaxl/synack-registry/pkg/users/userscontracts"
)

//RegistryHandler handler
type RegistryHandler struct {
	service regscontract.RegsService
}

//Bind controller
func (h *RegistryHandler) Bind(router *echo.Echo) {
	router.POST("/reg", h.POST)
}

//POST User
func (h *RegistryHandler) POST(c echo.Context) error {
	ctx := c.Request().Context()
	authUser, _ := c.Get("authenticated_user").(userscontracts.User)
	reg := regscontract.Request{}
	if err := c.Bind(&reg); err != nil {
		return errors.WithMessage(apierrors.ErrBadRequest, err.Error())
	}
	reg.Username = authUser.Username

	if err := h.service.Reg(ctx, &reg); err != nil {
		return err
	}

	c.JSON(http.StatusOK, reg)
	return nil
}
