package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/withnic/echotest/web/app/route"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	approute.Init(e)
}
