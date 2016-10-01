package approute

import (
	"github.com/NorifumiKawamoto/echotest/web/app/controllers/home"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init for app
func Init(e *echo.Echo) {
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string) bool {
		if username == "admin" && password == "secret" {
			return true
		}
		return false
	}))
	e.Static("/static", "static")
	e.GET("/home", homecontroller.Index)
}
