package home

import (
	"github.com/labstack/echo"
)

// Index home page default
func Index(c echo.Context) error {
	return HelloView(c)
}
