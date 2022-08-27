package handler

import (
	"github.com/labstack/echo/v4"
)

const (
	getMethod int = iota
	postMethod
)

type Handler struct {
	method      int
	Uri         string
	HandlerFunc echo.HandlerFunc
}

func GetHandlers() []Handler {
	var handlers []Handler

	for URL, FUNC := range getMapping() {
		handler := new(Handler)
		handler.method = getMethod
		handler.Uri = BASE_URL + URL
		handler.HandlerFunc = FUNC
		handlers = append(handlers, *handler)
	}
	for URL, FUNC := range postMapping() {
		handler := new(Handler)
		handler.method = postMethod
		handler.Uri = BASE_URL + URL
		handler.HandlerFunc = FUNC
		handlers = append(handlers, *handler)
	}

	return handlers
}

func (handler Handler) IsGetMethod() bool {
	if handler.method == getMethod {
		return true
	} else {
		return false
	}
}

func (handler Handler) IsPostMethod() bool {
	if handler.method == postMethod {
		return true
	} else {
		return false
	}
}

// func GetProjectLists() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		plArray := controller.GetProjectList()

// 		return c.JSON(http.StatusOK, plArray)
// 	}
// }
