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
var GET_PROJECT_URL = "project/:id"
var GET_SELECTABLE_MEMBER_LIST_URL = "selectablememberlist"
var GET_SELECTABLE_CLIENT_LIST_URL = "selectableclientlist"

// FUNCTION define
var GET_PROJECT_LIST_FUNC = func(c echo.Context) error {
	plArray := controller.GetProjectList()
	return c.JSON(http.StatusOK, plArray)
}
var GET_PROJECT_FUNC = func(c echo.Context) error {
	id := c.Param("id")
	p := controller.GetProjectItem(id)
	return c.JSON(http.StatusOK, p)
}
var GET_SELECTABLE_MEMBER_LIST_FUNC = func(c echo.Context) error {
	members := controller.GetSelectableMembers()
	return c.JSON(http.StatusOK, members)
}
var GET_SELECTABLE_CLIENT_LIST_FUNC = func(c echo.Context) error {
	clients := controller.GetSelectableClients()
	return c.JSON(http.StatusOK, clients)
}

// URL FUNCTION Mapping
func mapping() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		GET_PROJECT_LIST_URL:           GET_PROJECT_LIST_FUNC,
		GET_PROJECT_URL:                GET_PROJECT_FUNC,
		GET_SELECTABLE_MEMBER_LIST_URL: GET_SELECTABLE_MEMBER_LIST_FUNC,
		GET_SELECTABLE_CLIENT_LIST_URL: GET_SELECTABLE_CLIENT_LIST_FUNC,
	}
}
