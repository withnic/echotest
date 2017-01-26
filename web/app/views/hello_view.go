package views

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
)

// template helper
func HelloView(c echo.Context, r models.RenderObject) error {
	return c.Render(http.StatusOK, "helloindex", r)
}
