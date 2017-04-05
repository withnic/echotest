package views

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
)

type UserRender struct {
	u   models.User
	mes []models.Message
}

// UsersView is Users View Page Helper
func UsersView(c echo.Context, u []models.User) error {

	return c.Render(http.StatusOK, "userindex", u)
}

// UserEditView is User View Page Helper
func UserEditView(c echo.Context, u models.User) error {
	return c.Render(http.StatusOK, "useredit", u)
}

// UserView is User View Page Helper
func UserView(c echo.Context, u models.User) error {
	dest := make(map[string]interface{})
	dest["u"] = u
	dest["mes"] = u.Messages()

	return c.Render(http.StatusOK, "usershow", dest)
}

// UserFormView is User create Form Page Helper
func UserFormView(c echo.Context) error {
	return c.Render(http.StatusOK, "usernew", nil)
}
