package handler

import (
	"logicApi/src/main/net/baseonlura/slorder/controller"

	"net/http"

	"github.com/labstack/echo/v4"
)

// BASE_URL
var BASE_URL = "/api/"

// URL define
var GET_PROJECT_LIST_URL = "projectlist"

// FUNCTION define
var GET_PROJECT_LIST_FUNC = func(c echo.Context) error {
	plArray := controller.GetProjectList()
	return c.JSON(http.StatusOK, plArray)
}

// URL FUNCTION Mapping
func mapping() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		GET_PROJECT_LIST_URL: GET_PROJECT_LIST_FUNC,
	}
}
