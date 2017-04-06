package views

import (
	"net/http"

	"github.com/labstack/echo"
)

// HomeView is root
func HomeView(c echo.Context) error {
	dest := make(map[string]interface{})
	dest["Title"] = "Home"
	return c.Render(http.StatusOK, "homeindex", dest)
}
