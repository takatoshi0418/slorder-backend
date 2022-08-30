package handler

import (
	"logicApi/src/main/net/baseonlura/slorder/controller"

	"logicApi/src/main/net/baseonlura/slorder/viewModel"
	"strconv"

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
var GET_PROJECT_BELONG_WORK_LIST_URL = "projectbelongworklist/:id/:date"
var GET_SIMPLE_PROJECT_URL = "simpleproject/:id"
var SET_WORK_URL = "setwork"
var GET_PURE_PROJECT_URL = "getpureproject"
var SET_PROJECT_URL = "setproject"
var GET_SELECTABLE_OTHER_COST_KIND_LIST_URL = "selectableothercostkindlist"

// FUNCTION define
var GET_PROJECT_LIST_FUNC = func(c echo.Context) error {
	plArray, err := controller.GetProjectList()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, plArray)
}
var GET_PROJECT_FUNC = func(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	p, err := controller.GetProjectItem(id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, p)
}
var GET_SELECTABLE_MEMBER_LIST_FUNC = func(c echo.Context) error {
	members, err := controller.GetSelectableMembers()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, members)
}
var GET_SELECTABLE_CLIENT_LIST_FUNC = func(c echo.Context) error {
	clients, err := controller.GetSelectableClients()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, clients)
}

var GET_SIMPLE_PROJECT_FUNC = func(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	project, err := controller.GetSimpleProjectItem(id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, project)
}

var GET_PROJECT_BELONG_WORK_LIST_FUNC = func(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	dateStr := c.Param("date")

	works, err := controller.GetProjectMemberList(id, dateStr)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, works)
}

var SET_WORK_FUNC = func(c echo.Context) error {
	var pmVModels []viewModel.ProjectMember
	err := c.Bind(&pmVModels)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = controller.SetWorksByProjectMember(pmVModels)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "")
}

var GET_PURE_PROJECT_FUNC = func(c echo.Context) error {
	project, err := controller.GetPureProject()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, project)
}

var SET_PROJECT_FUNC = func(c echo.Context) error {

	var vModel viewModel.ProjectItem
	err := c.Bind(&vModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = controller.SetProjectByProjectItem(vModel)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "")
}

var GET_SELECTABLE_OTHER_COST_KIND_LIST_FUNC = func(c echo.Context) error {
	kinds, err := controller.GetOtherCostKinds()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, kinds)
}

// URL FUNCTION Mapping
func getMapping() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		GET_PROJECT_LIST_URL:                    GET_PROJECT_LIST_FUNC,
		GET_PROJECT_URL:                         GET_PROJECT_FUNC,
		GET_SELECTABLE_MEMBER_LIST_URL:          GET_SELECTABLE_MEMBER_LIST_FUNC,
		GET_SELECTABLE_CLIENT_LIST_URL:          GET_SELECTABLE_CLIENT_LIST_FUNC,
		GET_SIMPLE_PROJECT_URL:                  GET_SIMPLE_PROJECT_FUNC,
		GET_PROJECT_BELONG_WORK_LIST_URL:        GET_PROJECT_BELONG_WORK_LIST_FUNC,
		GET_PURE_PROJECT_URL:                    GET_PURE_PROJECT_FUNC,
		GET_SELECTABLE_OTHER_COST_KIND_LIST_URL: GET_SELECTABLE_OTHER_COST_KIND_LIST_FUNC,
	}
}

func postMapping() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		SET_WORK_URL:    SET_WORK_FUNC,
		SET_PROJECT_URL: SET_PROJECT_FUNC,
	}
}
