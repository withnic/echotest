package user

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

// Edit is show edit form
func Edit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(id)
	user := models.User{
		ID: id,
	}
	err = user.Get()

	return views.UserEditView(c, user)
}

// Update is update user
func Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(id)
	u := models.User{
		ID: id,
	}
	err = u.Get()
	if err != nil {
		log.Fatal(err)
	}
	u.Email = c.FormValue("email")
	u.Passwd = c.FormValue("passwd")

	if err := u.Validate(); err != nil {
		log.Fatal(err)
		return c.String(http.StatusOK, "NG")
	}

	err = u.Update()
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, u.ID)
}

// New is User create Form Page
func New(c echo.Context) error {
	return views.UserFormView(c)
}

// Create is create action
func Create(c echo.Context) error {
	u := models.User{}
	u.Email = c.FormValue("email")
	u.Passwd = c.FormValue("passwd")

	if err := u.Validate(); err != nil {
		log.Fatal(err)
		return c.String(http.StatusOK, "NG")
	}

	if err := u.Create(); err != nil {
		log.Fatal(err)
	}
	return c.Redirect(301, "/users")
}

// Delete is delete action
func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		ID: id,
	}
	err = user.Get()
	if err != nil {
		log.Fatal(err)
	}

	err = user.Delete()
	if err != nil {
		log.Fatal(err)
	}

	return c.NoContent(http.StatusNoContent)
}
