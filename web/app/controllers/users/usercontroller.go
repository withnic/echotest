package usercontroller

import (
	"fmt"
	"log"
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
		Db: models.UserDb{},
	}
	user = user.Get(id)
	return views.UserView(c, user)
}

// New is User create Form Page
func New(c echo.Context) error {
	//	return views.UserNew(c)
	return nil
}

func Create(u models.User) error {
	// u := models.User{
	// 	Id:    3,
	// 	Email: "test@test.com",
	// }
	return nil
	//	return u.Create()
}
