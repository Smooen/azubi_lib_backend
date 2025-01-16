package main

import (
	"log"
	"net/http"
	"os"

	"azubi_library/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	os.Remove("./test.sqlite")
	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	handler := controllers.Handler{DB: db}

	migrate(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	SetupRoutes(e, &handler)

	e.Logger.Fatal(e.Start(":1323"))
}
