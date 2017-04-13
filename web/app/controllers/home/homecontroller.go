package home

import (
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
)

// Index home page default
func Index(c echo.Context) error {
	//	mu := models.MessageWithUser{}
	//	messagesWithUser := mu.GetAll()
	u := models.User{
		ID: 1,
	}

	messagesWithUser := u.GetAllFollowersMessage()

	return views.HomeView(c, messagesWithUser)
}
