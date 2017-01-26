package approute

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/withnic/echotest/web/app/controllers/home"
	"github.com/withnic/echotest/web/app/controllers/users"
	"github.com/withnic/echotest/web/app/libs"
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
	e.GET("/", homecontroller.Index)
	e.GET("/users", usercontroller.Index)
}
