package main

import (
	"logicApi/src/main/net/baseonlura/slorder/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// create echo instance
	e := echo.New()

	// define middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.GetOneProject())

	e.Logger.Fatal(e.Start(":1323"))
}
