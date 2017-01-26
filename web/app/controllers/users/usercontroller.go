package usercontroller

import (
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
)

// Index home page default
func Index(c echo.Context) error {
	user := models.User{}
	users := user.GetAll()
	return views.UsersView(c, users)
}
