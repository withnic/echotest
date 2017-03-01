package usercontroller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// Show home page default
func Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(id)
	user := models.User{
		ID: id,
	}
	err = user.Get()
	if err != nil {
		log.Fatal(err)
	}
	return views.UserView(c, user)
}

// New is User create Form Page
func New(c echo.Context) error {
	return views.UserFormView(c)
}

// Create is create action
func Create(c echo.Context) error {
	u := models.User{}
	email := c.FormValue("email")
	u.Email = email

	if err := u.Validate(); err != nil {
		log.Fatal(err)
		return c.String(http.StatusOK, "NG")
	}

	if err := u.Create(); err != nil {
		log.Fatal(err)
	}
	return c.String(http.StatusOK, "OK")
}
