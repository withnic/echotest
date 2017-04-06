package auth

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
	"github.com/withnic/echotest/web/app/views"
)

func Login(c echo.Context) error {
	return views.LoginView(c)
}

func Signin(c echo.Context) error {
	u := models.User{}
	u.Email = c.FormValue("email")
	error := u.GetByEmail()

	if error != nil {
		log.Fatal(error)
	}
	password := []byte(c.FormValue("passwd"))
	error = bcrypt.CompareHashAndPassword([]byte(u.Passwd), password)

	if error != nil {
		log.Fatal(error)
	}

	return c.JSON(http.StatusOK, "OK")
}
