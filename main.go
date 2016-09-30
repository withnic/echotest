package main

import (
	"github.com/NorifumiKawamoto/echotest/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.Get("/hello", handler.MainPage())

	// サーバー起動
	e.Run(standard.New(":8080"))
}
