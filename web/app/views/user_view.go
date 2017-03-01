package views

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
)

// UsersView is Users View Page Helper
func UsersView(c echo.Context, r []models.User) error {

	return c.Render(http.StatusOK, "usersindex", r)
}

// UserView is User View Page Helper
func UserView(c echo.Context, r models.User) error {
	return c.Render(http.StatusOK, "usershow", r)
}

// UserFormView is User create Form Page Helper
func UserFormView(c echo.Context) error {
	return c.Render(http.StatusOK, "usernew", nil)
}
