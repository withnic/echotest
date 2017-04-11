package views

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
)

func current(c echo.Context) (models.User, error) {
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

		return u, nil
	}

	return models.User{}, errors.New("not found")
}

func deleteCookie(c echo.Context, name string) {
	cookie, _ := c.Cookie(name)
	if cookie != nil {
		cookie.Value = ""
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		c.SetCookie(cookie)
	}
}
