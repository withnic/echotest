package home

import (
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/views"
)

// Index home page default
func Index(c echo.Context) error {
	return views.HomeView(c)
}
