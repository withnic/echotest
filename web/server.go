package main

import (
	"github.com/NorifumiKawamoto/echotest/web/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Debug()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	routes.Init(e)

	// サーバー起動
	e.Run(standard.New(":8080"))
}
