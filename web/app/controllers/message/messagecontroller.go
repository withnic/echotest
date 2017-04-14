package message

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
)

// New is User create Form Page
func New(c echo.Context) error {
	return views.MessagesFormView(c)
}

// Create is create action
func Create(c echo.Context) error {
	mes := models.Message{}
	mes.UserID, _ = getUserID(c)
	mes.Body = c.FormValue("body")
	if err := mes.Validate(); err != nil {
		setSession(c, "errormessage", "Invalid params.")
		return views.UserFormView(c)
	}

	if err := mes.Create(); err != nil {
		setSession(c, "errormessage", "Invalid params.")
		return views.UserFormView(c)
	}

	return c.Redirect(301, "/")
}

//本来セッションに入れる（echoのsessionが動かないので今回はやらない)
func setSession(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Path = "/"
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

// getUserID is return User.id
func getUserID(c echo.Context) (int, error) {
	cookie, _ := c.Cookie("uid")
	if cookie != nil {
		i, _ := strconv.Atoi(cookie.Value)
		u := models.User{
			ID: i,
		}
		err := u.Get()
		if err != nil {
			log.Fatal(err)
		}
		return u.ID, nil
	}
	return 0, errors.New("not found")
}
