package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Books struct {
	gorm.Model
	Title  string
	Isbn string
	Author string
	ReleaseDate string
	Availability bool
}

var books =[]Book { 
	{Isbn: "9780441172719", Title: "Dune", Author: "Frank Herbert", ReleaseDate: "1987", Availability: false}, 
}

func getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books);
}

func main() {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	//db, err := sql.Open("sqlite3", "./test.sqlite")
	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&Books{
		Title:  "Test",
		Isbn: "91203001",
		Author: "Ich habs geschrieben",
		//ReleaseDate: "2024",
		Availability: true,
	})

	//createBooksTable(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.GET("/Books", getBooks)

	e.Logger.Fatal(e.Start(":1323"))
	
}

type Book struct {
	Title  string `json:"title" query:"title"`
	Isbn string `json:"isbn" query:"isbn"`
	Author string `json:"author" query:"author"`
	ReleaseDate string `json:"releaseDate" query:"releaseDate"`
	Availability bool `json:"availability" query:"availability"`
}
