package user

import (
	"net/http"

	"github.com/labstack/echo"
)

// UsersView is Users View Page Helper
func UsersView(c echo.Context, r []User) error {

	return c.Render(http.StatusOK, "userindex", r)
}

// UserEditView is User View Page Helper
func UserEditView(c echo.Context, r User) error {
	return c.Render(http.StatusOK, "useredit", r)
}

// UserView is User View Page Helper
func UserView(c echo.Context, r User) error {
	return c.Render(http.StatusOK, "usershow", r)
}

// UserFormView is User create Form Page Helper
func UserFormView(c echo.Context) error {
	return c.Render(http.StatusOK, "usernew", nil)
}
