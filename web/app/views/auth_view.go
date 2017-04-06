package views

import (
	"net/http"

	"github.com/labstack/echo"
)

// LoginView is root
func LoginView(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}
