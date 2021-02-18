package api

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/kataras/golog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	sl "github.com/mattn/go-sqlite3"
	"github.com/skyaxl/synack-registry/pkg/regs/regsservice"
	userservice "github.com/skyaxl/synack-registry/pkg/users/usersservice"
)

//New echo api
func New() (err error) {
	var db *sql.DB
	golog.Info(sl.Version())
	if db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?_loc=UTC", "./db/prod.db")); err != nil {
		return
	}
	svc := userservice.New(db)
	regService := regsservice.New(db)
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		c.String(200, "PONG")
		return nil
	})
	authMid := AuthHandler{svc}
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: authMid.Middleware,
		Skipper: func(c echo.Context) bool {
			if c.Request().URL.Path == "/ping" {
				return true
			}
			return false
		},
	}))
	userHandler := &UsersHandler{svc}
	userHandler.Bind(e)

	regHandler := &RegistryHandler{regService}
	regHandler.Bind(e)
	portStr := os.Getenv("APP_PORT")
	port, _ := strconv.ParseInt(portStr, 10, 64)
	if port == 0 {
		port = 8080
	}
	err = e.Start(fmt.Sprintf(":%d", port))
	golog.Infof("[Registry Api] Has started with address localhost:%d\n", port)
	return
}
