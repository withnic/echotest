package home

import (
	"net/http"

	"github.com/labstack/echo"
)

// template helper
func HelloView(c echo.Context) error {
	var Title interface{} = "Hello"
	return c.Render(http.StatusOK, "homeindex", Title)
}
