package home

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
)

// Index home page default
func Index(c echo.Context) error {
	var messagesWithUser []models.MessageWithUser
	uid, err := getSession(c, "uid")
	if err != nil {
		mu := models.MessageWithUser{}
		messagesWithUser = mu.GetAll()
	} else {
		i, _ := strconv.Atoi(uid)
		u := models.User{
			ID: i,
		}
		messagesWithUser = u.GetAllFollowersAndMeMessage()
	}

	return views.HomeView(c, messagesWithUser)
}

//本来セッション取得（echoのsessionが動かないので今回はやらない)
func getSession(c echo.Context, name string) (string, error) {
	cookie, _ := c.Cookie(name)
	if cookie != nil {
		return cookie.Value, nil
	}

	return "", errors.New("not found")
}
