package views

import (
	"net/http"

	"github.com/NorifumiKawamoto/echotest/web/app/models"
	"github.com/labstack/echo"
)

// template helper
func UsersView(c echo.Context, r []models.User) error {

	return c.Render(http.StatusOK, "usersindex", r)
}
