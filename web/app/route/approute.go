package approute

import (
	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/controllers/auth"
	"github.com/withnic/echotest/web/app/controllers/home"
	"github.com/withnic/echotest/web/app/controllers/message"
	"github.com/withnic/echotest/web/app/controllers/user"
	"github.com/withnic/echotest/web/app/libs"
)

// Init for app
func Init(e *echo.Echo) {
	libs.SetRenderer(e, "app/template/**/*.html")
	e.Static("/static", "static")
	e.GET("/", home.Index)
	e.GET("/users/:id/follow", user.Follow)
	e.GET("/users/:id/unfollow", user.UnFollow)

	e.GET("/users/:id/edit", user.Edit)
	e.GET("/users/:id", user.Show)
	e.PUT("/users/:id", user.Update)
	e.DELETE("/users/:id", user.Delete)

	e.GET("/users", user.Index)

	e.POST("/messages/create", message.Create)

	e.GET("/signup", user.New)
	e.POST("/users", user.Create)
	e.GET("/login", auth.Login)
	e.GET("/logout", auth.Logout)
	e.POST("/signin", auth.Signin)

}
