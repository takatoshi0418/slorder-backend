package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"logicApi/src/main/net/baseonlura/slorder/handler"
)

func main() {
	// create echo instance
	e := echo.New()

	// define middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	for _, handler := range handler.GetHandlers() {
		e.GET(handler.Uri, handler.HandlerFunc)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
