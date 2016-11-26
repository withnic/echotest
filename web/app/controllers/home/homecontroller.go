package homecontroller

import (
	"github.com/NorifumiKawamoto/echotest/web/app/models"
	"github.com/NorifumiKawamoto/echotest/web/app/views"
	"github.com/labstack/echo"
)

// Index home page default
func Index(c echo.Context) error {
	data := models.RenderObject{
		Title: "Hello, world",
		Name:  "Echo-Test",
		Lang:  "hoge",
	}

	return views.HelloView(c, data)
}
