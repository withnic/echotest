package home

import (
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
)

// Index home page default
func Index(c echo.Context) error {
	mes := models.Message{}
	messages := mes.GetAll()
	return views.HomeView(c, messages)
}
