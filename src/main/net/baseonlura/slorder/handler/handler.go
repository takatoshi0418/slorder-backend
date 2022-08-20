package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Method      string
	Uri         string
	HandlerFunc echo.HandlerFunc
}

func GetHandlers() []Handler {
	var handlers []Handler

	for URL, FUNC := range mapping() {
		handler := new(Handler)
		handler.Uri = BASE_URL + URL
		handler.HandlerFunc = FUNC
		handlers = append(handlers, *handler)
	}

	return handlers
}

// func GetProjectLists() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		plArray := controller.GetProjectList()

// 		return c.JSON(http.StatusOK, plArray)
// 	}
// }
