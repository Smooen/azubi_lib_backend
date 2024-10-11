package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type Book struct {
	Title  string `json:"title" query:"title"`
	Isbn string `json:"isbn" query:"isbn"`
	Author string `json:"author" query:"author"`
	ReleaseDate string `json:"releaseDate" query:"releaseDate"`
	Availability bool string `json:"availability" query:"availability"`
}