package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ebi-yade/go-echo-test/config"
	"github.com/ebi-yade/go-echo-test/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.RootHandler())

	e.Start(":" + config.Port())
}
