package homecontroller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index home page default
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world")
}
