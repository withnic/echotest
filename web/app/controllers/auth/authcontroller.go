package auth

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
)

var store = sessions.NewCookieStore([]byte("random-string"))

func Login(c echo.Context) error {
	return views.LoginView(c)
}

func Signin(c echo.Context) error {
	u := models.User{}
	u.Email = c.FormValue("email")
	error := u.GetByEmail()

	if error != nil {
		errorMessage(c)
		return views.LoginView(c)
	}

	password := []byte(c.FormValue("passwd"))
	error = bcrypt.CompareHashAndPassword([]byte(u.Passwd), password)

	if error != nil {
		fmt.Println(error.Error())
		errorMessage(c)
		return views.LoginView(c)
	}

	return c.JSON(http.StatusOK, "OK")
}

//本来セッションに入れる（echoのsessionが動かないので今回はやらない)
func errorMessage(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "errormessage"
	cookie.Path = "/"
	cookie.Value = "not exists email or password wrong."
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}
