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

// Login is view login form
func Login(c echo.Context) error {
	return views.LoginView(c, http.StatusOK)
}

// Signin is add User session
func Signin(c echo.Context) error {
	u := models.User{}
	u.Email = c.FormValue("email")
	error := u.GetByEmail()
	if error != nil || u.ID == 0 {
		setSession(c, "errormessage", "not exists email or password wrong.")
		//  httpjarを使わないとredirectでset-cookie消える
		//  一旦スルー
		return views.LoginView(c, http.StatusBadRequest)
	}

	password := []byte(c.FormValue("passwd"))
	error = bcrypt.CompareHashAndPassword([]byte(u.Passwd), password)
	if error != nil {
		setSession(c, "errormessage", "not exists email or password wrong.")
		//  httpjarを使わないとredirectでset-cookie消える
		//  一旦スルー
		return views.LoginView(c, http.StatusBadRequest)
	}
	setSession(c, "uid", fmt.Sprint(u.ID))
	return c.JSON(http.StatusOK, "OK")
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
