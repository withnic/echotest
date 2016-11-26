package views

import (
	"net/http"

	"github.com/NorifumiKawamoto/echotest/web/app/models"
	"github.com/labstack/echo"
)

// template helper
func HelloView(c echo.Context, r models.RenderObject) error {
	return c.Render(http.StatusOK, "index", r)
}
