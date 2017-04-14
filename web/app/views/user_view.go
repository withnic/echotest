package views

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
)

// UsersView is Users View Page Helper
func UsersView(c echo.Context, u []models.User) error {
	dest := make(map[string]interface{})
	dest["u"] = u
	dest["Title"] = "User"

	current, err := current(c)
	if err == nil {
		dest["Current"] = current
	}
	return c.Render(http.StatusOK, "userindex", dest)
}

// UserEditView is User View Page Helper
func UserEditView(c echo.Context, u models.User) error {
	dest := make(map[string]interface{})
	dest["u"] = u
	dest["Title"] = "User"

	current, err := current(c)
	if err == nil {
		dest["Current"] = current
	}

	return c.Render(http.StatusOK, "useredit", dest)
}

// UserView is User View Page Helper
func UserView(c echo.Context, u models.User) error {
	dest := make(map[string]interface{})
	dest["u"] = u
	dest["mes"] = u.Messages()
	dest["Title"] = "User"

	current, err := current(c)
	if err == nil {
		dest["Current"] = current

		dest["IsFollow"] = current.IsFollow(u.ID)
	}
	return c.Render(http.StatusOK, "usershow", dest)
}

// UserFormView is User create Form Page Helper
func UserFormView(c echo.Context) error {
	dest := make(map[string]interface{})
	dest["Title"] = "User"
	current, err := current(c)
	if err == nil {
		dest["Current"] = current
	}
	return c.Render(http.StatusOK, "usernew", dest)
}
