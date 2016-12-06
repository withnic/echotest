package usercontroller

import (
	"github.com/NorifumiKawamoto/echotest/web/app/models"
	"github.com/NorifumiKawamoto/echotest/web/app/views"
	"github.com/labstack/echo"
)

// Index home page default
func Index(c echo.Context) error {
	user := models.User{}
	users := user.GetAll()
	return views.UsersView(c, users)
}
