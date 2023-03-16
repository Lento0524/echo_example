package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	c := newController(db)
	e.GET("/food/:id", c.get)
	e.GET("/food", getAll)
	e.POST("/food", post)
	e.PATCH("/food/:id", patch)
	e.DELETE("/food/:id", delete)
	e.Logger.Fatal(e.Start(":1323"))
}

type foodGetRequest struct {
	Id int `param:"id"`
}

type foodPostRequest struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type foodPatchRequest struct {
	Id   int    `param:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type foodDeleteRequest struct {
	Id int `param:"id"`
}

type foodResponce struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type controller struct {
	db *sql.DB
}

func newController(db *sql.DB) controller {
	return controller{
		db: db,
	}
}

func (a controller) get(c echo.Context) error {

	req := new(foodGetRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, foodResponce{
		Id:        req.Id,
		Name:      "",
		Unit:      "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
}

func getAll(c echo.Context) error {
	resAll := []foodResponce{
		{
			Id:        0,
			Name:      "ああ",
			Unit:      "",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			Id:        1,
			Name:      "いい",
			Unit:      "",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
	return c.JSON(http.StatusOK, resAll)
}

func post(c echo.Context) error {
	req := new(foodPostRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, foodResponce{
		Id:        0,
		Name:      req.Name,
		Unit:      req.Unit,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
}

func patch(c echo.Context) error {
	req := new(foodPatchRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, foodResponce{
		Id:        req.Id,
		Name:      req.Name,
		Unit:      req.Unit,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
}

func delete(c echo.Context) error {
	req := new(foodDeleteRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, req)
}
