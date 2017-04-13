package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/withnic/echotest/web/routes"
)

func main() {
	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	routes.Init(e)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8081"))
}
