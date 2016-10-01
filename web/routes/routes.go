package routes

import (
	"github.com/NorifumiKawamoto/echotest/web/app/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	approute.Init(e)
}
