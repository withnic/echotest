package views

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
)

// template helper
func UsersView(c echo.Context, r []models.User) error {

	return c.Render(http.StatusOK, "usersindex", r)
}
