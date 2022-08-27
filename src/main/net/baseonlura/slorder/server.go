package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"logicApi/src/main/net/baseonlura/slorder/handler"

	"logicApi/src/main/net/baseonlura/slorder/db"
)

func main() {
	// create echo instance
	e := echo.New()

	// define middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	for _, handler := range handler.GetHandlers() {
		if handler.IsGetMethod() {
			e.GET(handler.Uri, handler.HandlerFunc)
		} else if handler.IsPostMethod() {
			e.POST(handler.Uri, handler.HandlerFunc)
		}
	}

	//connection connection
	connection := new(db.DBConnect)
	err := connection.Initializer()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.SetDBConnection(connection)

	e.Logger.Fatal(e.Start(":1323"))
}
