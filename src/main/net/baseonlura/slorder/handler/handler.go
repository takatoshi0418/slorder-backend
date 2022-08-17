package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var project string = "{\n \"project\": {\n \"status\": 1,\n \"basic\": {\n \"no\": \"P-20220301-0002\",\n \"name\": \"ペット行動管理システム\",\n \"client\": 2,\n \"startDate\": \"2022-03-01\",\n \"limitDate\": \"2022-04-30\",\n \"receiveAmount\": 20000000\n },\n \"payment\": {\n \"estimate\": {\n \"oparatingWorkByTime\": 10,\n \"oparatingCost\": 8640000,\n \"otherCost\": 5400000,\n \"proceeds\": 5496000\n },\n \"actual\": {\n \"oparatingWorkByTime\": 3.75,\n \"oparatingCost\": 1435735,\n \"otherCost\": 5400000,\n \"proceeds\": 13164265\n }\n },\n \"members\": [\n {\n \"value\": 1,\n \"unit\": 2500,\n \"oparatingTime\": 150\n },\n {\n \"value\": 2,\n \"unit\": 4000,\n \"oparatingTime\": 90\n },\n {\n \"value\": 3,\n \"unit\": 3500,\n \"oparatingTime\": 100.01\n },\n {\n \"value\": 4,\n \"unit\": 3500,\n \"oparatingTime\": 120.4\n }\n ],\n \"otherCosts\": [\n {\n \"name\": \"サーバー01\",\n \"kind\": 1,\n \"buyDate\": \"2022-03-01\",\n \"price\": 5000000\n },\n {\n \"name\": \"武田信玄\",\n \"kind\": 2,\n \"buyDate\": \"2022-03-04\",\n \"price\": 400000\n }\n ],\n \"histories\": [\n {\n \"name\": \"受注 太郎\",\n \"date\": \"2022-02-01\",\n \"kind\": 0\n },\n {\n \"name\": \"受注 太郎\",\n \"date\": \"2022-02-20\",\n \"kind\": 1\n }\n ]\n }\n}"

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	}
}

func GetOneProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, project)
	}
}
