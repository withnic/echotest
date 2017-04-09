package approute

import (
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/controllers/auth"
	"github.com/withnic/echotest/web/app/controllers/home"
	"github.com/withnic/echotest/web/app/controllers/user"
	"github.com/withnic/echotest/web/app/libs"
)

// Init for app
func Init(e *echo.Echo) {
	libs.SetRenderer(e, "app/template/**/*.html")
	e.Static("/static", "static")
	e.GET("/", home.Index)
	e.GET("/users", user.Index)
	e.GET("/users/:id", user.Show)
	e.GET("/users/:id/edit", user.Edit)
	e.PUT("/users/:id", user.Update)
	e.DELETE("/users/:id", user.Delete)
	e.GET("/users/new", user.New)
	e.POST("/users", user.Create)
	e.GET("/login", auth.Login)
	e.POST("/signin", auth.Signin)

}
