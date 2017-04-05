package views

import (
	"net/http"

	"github.com/labstack/echo"
)

// HomeView is root
func HomeView(c echo.Context) error {
	var Title interface{} = "Hello"
	return c.Render(http.StatusOK, "homeindex", Title)
}
