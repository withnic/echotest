package homecontroller

import (
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
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
