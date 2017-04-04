package approute

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/withnic/echotest/web/app/home"
	"github.com/withnic/echotest/web/app/libs"
	user "github.com/withnic/echotest/web/app/users"
)

// Init for app
func Init(e *echo.Echo) {
	libs.SetRenderer(e, "app/template/**/*.html")
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string) bool {
		if username == "admin" && password == "secret" {
			return true
		}
		return false
	}))
	e.Static("/static", "static")
	e.GET("/", home.Index)
	e.GET("/users", user.Index)
	e.GET("/users/:id", user.Show)
	e.GET("/users/new", user.New)
	e.POST("/users", user.Create)
}
