package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/food", post)
	e.PATCH("/food/:id", patch)
	e.DELETE("/food/:id", delete)
	e.Logger.Fatal(e.Start(":1323"))
}

type foodPostRequest struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type foodPatchRequest struct {
	Id   string `param:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type foodDeleteRequest struct {
	Id string `param:"id"`
}

func post(c echo.Context) error {
	req := new(foodPostRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, req)
}

func patch(c echo.Context) error {
	req := new(foodPatchRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, req)
}

func delete(c echo.Context) error {
	req := new(foodDeleteRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, req)
}
